package cgtn

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) GetTop(req *ghttp.Request) {

	if value, err := g.Redis().DoVar("GET", "CGTN_TOP"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.cgtn.com/"
	rssData := dao.RSSFeed{
		Title:    "CGTN - Top News",
		Link:     apiUrl,
		Tag:      []string{"英文", "海外"},
		ImageUrl: "https://ui.cgtn.com/static/ng/resource/images/logo_title.png",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		docs := soup.HTMLParse(resp.ReadAllString())
		topItems := docs.FindAll("div", "class", "top-news-item")
		rssItems := make([]dao.RSSItem, 0)
		for _, item := range topItems {
			rssItem := dao.RSSItem{}
			contentHtml := item.Find("div", "class", "top-news-item-content")
			title := contentHtml.Find("a").Text()
			time := contentHtml.Find("a").Attrs()["data-time"]
			link := contentHtml.Find("a").Attrs()["href"]
			description := getMainContent(link)

			rssItem.Title = title
			rssItem.Link = link
			rssItem.Created = time
			rssItem.Description = description
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := lib.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "36KR_NEWS_FLASHES", rssStr)
	g.Redis().DoVar("EXPIRE", "36KR_NEWS_FLASHES", 60*60*1)
	_ = req.Response.WriteXmlExit(rssStr)
}
