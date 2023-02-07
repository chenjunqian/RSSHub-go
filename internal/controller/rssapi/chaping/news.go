package chaping

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetNews(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getNewsLinks()[linkType]

	cacheKey := "CHAPING_NEWS_" + linkConfig.Caty
	if value, err := service.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	targetUrl := "https://chaping.cn/news?cate=" + linkConfig.Caty
	currentUrl := "https://chaping.cn/api/official/information/news?page=1&limit=16&cate=" + linkConfig.Caty
	rssData := dao.RSSFeed{
		Title:       "差评-" + linkConfig.Title,
		Link:        targetUrl,
		Description: "新媒体「差评」站在科技领域，以专业的角度和诙谐的语言，为订阅者呈现高质量的图文和视频内容",
		ImageUrl:    "https://chaping.cn/public/favicon.ico",
	}
	if resp := service.GetContent(ctx,currentUrl); resp != "" {
		respJson := gjson.New(resp)
		dataJsonList := respJson.GetJsons("data")
		rssItems := make([]dao.RSSItem, 0)
		for _, dataJson := range dataJsonList {
			title := dataJson.Get("title").String()
			summary := dataJson.Get("summary").String()
			imageLink := dataJson.Get("images.url").String()
			time := dataJson.Get("time_publish_timestamp").String()
			author := dataJson.Get("author").String()
			link := dataJson.Get("origin_url").String()
			fullContent := dataJson.Get("content").String()
			rssItem := dao.RSSItem{
				Title:     title,
				Link:      link,
				Content:   feed.GenerateContent(summary + "<br>" + fullContent),
				Author:    author,
				Created:   time,
				Thumbnail: imageLink,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	service.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	service.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*3)
	req.Response.WriteXmlExit(rssStr)
}
