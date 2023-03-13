package baijing

import (
	"context"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetWeekly(req *ghttp.Request) {

	var ctx context.Context = context.Background()
	if value, err := cache.GetCache(ctx, "BAIJING_WEEKLY"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.baijingapp.com/index/ajax/getlist/type-1"
	rssData := dao.RSSFeed{
		Title:       "白鲸出海-7x24h",
		Link:        "https://www.baijingapp.com/",
		Description: "白鲸出海快讯",
		Tag:         []string{"新闻"},
		ImageUrl:    "https://www.baijingapp.com/static/css/default/img/favicon.ico",
	}
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		docs := soup.HTMLParse(resp)
		articleDocList := docs.FindAll("div", "id", "menuKuaixun")
		rssItems := make([]dao.RSSItem, 0)
		for _, articleDoc := range articleDocList {
			title := articleDoc.Find("h3").Text()
			time := articleDoc.Find("span").HTML()
			content := articleDoc.Find("div", "class", "newsflashesText").Text()

			rssItem := dao.RSSItem{
				Title:   title,
				Content: feed.GenerateContent(content),
				Author:  "",
				Created: time,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx, "BAIJING_WEEKLY", rssStr)
	req.Response.WriteXmlExit(rssStr)
}
