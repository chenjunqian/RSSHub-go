package dongqiudi

import (
	"context"
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetTopNews(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getNewsLinks()[linkType]

	cacheKey := "DONGQIUDI_TOP_NEWS_" + linkConfig.ChannelId
	if value, err := component.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := fmt.Sprintf("https://api.dongqiudi.com/app/tabs/iphone/%s.json?mark=gif&version=576", linkConfig.ChannelId)
	rssData := dao.RSSFeed{
		Title:       "懂球帝 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         linkConfig.Tags,
		Description: "懂球帝|专业权威的足球网站",
		ImageUrl:    "https://static1.dongqiudi.com/web-new/web/images/fav.ico",
	}
	if resp := component.GetContent(ctx,apiUrl); resp != "" {
		respJson := gjson.New(resp)
		articleJsonList := respJson.GetJsons("articles")
		rssItems := make([]dao.RSSItem, 0)
		for _, articleJson := range articleJsonList {
			title := articleJson.Get("title").String()
			link := articleJson.Get("url").String()
			if !strings.HasPrefix(link, "http") {
				link = articleJson.Get("share").String()
			}
			imageLink := articleJson.Get("thumb").String()
			time := articleJson.Get("published_at").String()
			author := articleJson.Get("author_classify").String()
			rssItem := dao.RSSItem{
				Title:     title,
				Link:      link,
				Content:   feed.GenerateContent(imageLink),
				Author:    author,
				Created:   time,
				Thumbnail: imageLink,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	component.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*6)
	req.Response.WriteXmlExit(rssStr)
}
