package cgtn

import (
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *controller) GetMostRead(req *ghttp.Request) {
	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]

	if value, err := g.Redis().DoVar("GET", "CGTN_MOST_"+linkType); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.cgtn.com/most-read"
	rssData := dao.RSSFeed{
		Title:    "CGTN - Most " + linkType,
		Link:     apiUrl,
		Tag:      []string{"英文", "海外"},
		ImageUrl: "https://ui.cgtn.com/static/ng/resource/images/logo_title.png",
	}
	if resp := component.GetContent(apiUrl); resp != "" {
		docs := soup.HTMLParse(resp)
		itemWrapper := docs.Find("div", "id", linkType+"Items")
		items := itemWrapper.FindAll("div", "class", "most-read-item-box")
		rssItems := make([]dao.RSSItem, 0)
		for _, item := range items {
			var (
				title       string
				time        string
				link        string
				cgContent   string
				mainContent string
				content     string
			)
			var rssItem = dao.RSSItem{}
			title = item.Find("a").Text()
			time = item.Find("a").Attrs()["data-time"]
			link = item.Find("a").Attrs()["href"]
			cgContent = item.Find("div", "class", "cg-content").Text()
			mainContent = getMainContent(link)
			content = fmt.Sprintf("%s<br>%s", cgContent, mainContent)
			content = feed.GenerateContent(content)
			rssItem.Title = title
			rssItem.Link = link
			rssItem.Created = time
			rssItem.Content = content
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "36KR_NEWS_FLASHES", rssStr)
	g.Redis().DoVar("EXPIRE", "36KR_NEWS_FLASHES", 60*60*1)
	_ = req.Response.WriteXmlExit(rssStr)
}
