package bishijie

import (
	"context"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetShenDu(req *ghttp.Request) {

	var ctx context.Context = context.Background()
	if value, err := component.GetRedis().Do(ctx,"GET", "BISHIJIE_SHENDU"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.bishijie.com/shendu"
	rssData := dao.RSSFeed{
		Title:       "币世界-深度",
		Link:        "https://www.baijingapp.com",
		Tag:         []string{"比特币", "金融", "科技", "投资"},
		Description: "币世界网-比特币等数字货币交易所导航、投资理财、快讯、深度、币圈、市场行情第一站。",
		ImageUrl:    "https://www.bishijie.com/favicon.ico",
	}
	if resp := component.GetContent(ctx,apiUrl); resp != "" {
		docs := soup.HTMLParse(resp)
		articleDocsList := docs.FindAll("div", "class", "articles-card")
		rssItems := make([]dao.RSSItem, 0)
		for _, articleDoc := range articleDocsList {
			var link string
			var imageLink string
			articleTitleDoc := articleDoc.Find("h4", "class", "articles-title")
			var title string
			if articleTitleDoc.Error == nil {
				title = articleDoc.Find("h4", "class", "articles-title").Text()
				title = strings.ReplaceAll(title, "\n", "")
				title = strings.ReplaceAll(title, " ", "")
			}
			articleImgDoc := articleDoc.Find("div", "class", "articles-img")
			if articleImgDoc.Error == nil {
				link = "https://www.bishijie.com" + articleImgDoc.Find("a").Attrs()["href"]
				imageLink = articleImgDoc.Find("img").Attrs()["src"]
			}
			content := articleDoc.Find("p", "class", "articles-text").Text()
			rssItem := dao.RSSItem{
				Title:     title,
				Link:      link,
				Content:   feed.GenerateContent(content),
				Thumbnail: imageLink,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx,"SET", "BISHIJIE_SHENDU", rssStr)
	component.GetRedis().Do(ctx,"EXPIRE", "BISHIJIE_SHENDU", 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
