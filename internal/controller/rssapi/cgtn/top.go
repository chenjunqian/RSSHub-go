package cgtn

import (
	"context"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetTop(req *ghttp.Request) {

	var ctx context.Context = context.Background()
	if value, err := service.GetRedis().Do(ctx,"GET", "CGTN_TOP"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.cgtn.com/"
	rssData := dao.RSSFeed{
		Title:    "CGTN - Top News",
		Link:     apiUrl,
		Tag:      []string{"英文", "海外"},
		ImageUrl: "https://ui.cgtn.com/static/ng/resource/images/logo_title.png",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != "" {
		docs := soup.HTMLParse(resp)
		topItems := docs.FindAll("div", "class", "top-news-item")
		rssItems := make([]dao.RSSItem, 0)
		for _, item := range topItems {
			var (
				title   string
				time    string
				link    string
				content string
			)
			rssItem := dao.RSSItem{}
			contentHtml := item.Find("div", "class", "top-news-item-content")
			title = contentHtml.Find("a").Text()
			time = contentHtml.Find("a").Attrs()["data-time"]
			link = contentHtml.Find("a").Attrs()["href"]
			content = getMainContent(ctx, link)

			rssItem.Title = title
			rssItem.Link = link
			rssItem.Created = time
			rssItem.Content = content
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	service.GetRedis().Do(ctx,"SET", "36KR_NEWS_FLASHES", rssStr)
	service.GetRedis().Do(ctx,"EXPIRE", "36KR_NEWS_FLASHES", 60*60*1)
	req.Response.WriteXmlExit(rssStr)
}
