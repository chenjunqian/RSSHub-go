package sspai

import (
	"context"
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]
	cacheKey := "SSPAI_INDEX_" + linkConfig.ChannelId
	if value, err := component.GetRedis().Do(ctx, "GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	var apiUrl string
	if linkConfig.ChannelId == "recommend" {
		apiUrl = "https://sspai.com/api/v1/article/index/page/get?limit=10&offset=0&created_at=0"
	} else {
		apiUrl = fmt.Sprintf("https://sspai.com/api/v1/article/tag/page/get?limit=10&offset=0&tag=%s", linkConfig.Title)
	}
	rssData := dao.RSSFeed{
		Title:       "少数派 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         []string{"科技"},
		Description: "少数派致力于更好地运用数字产品或科学方法，帮助用户提升工作效率和生活品质",
		ImageUrl:    "https://cdn.sspai.com/sspai/assets/img/favicon/icon.ico",
	}

	if resp := component.GetContent(ctx, apiUrl); resp != "" {

		rssItems := commonParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx, "SET", cacheKey, rssStr)
	component.GetRedis().Do(ctx, "EXPIRE", cacheKey, 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
