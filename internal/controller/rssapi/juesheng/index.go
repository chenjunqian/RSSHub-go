package juesheng

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetIndex(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]
	cacheKey := "JUESHENG_INDEX_" + linkConfig.ChannelUrl
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := linkConfig.ChannelUrl
	rssData := dao.RSSFeed{
		Title:       "决胜网 - " + linkConfig.Title,
		Link:        apiUrl,
		Description: "决胜网是教育产业门户网站提供：教育门户新闻资讯、互联网+教育、在线教育、兴趣教育、在线职业教育、教育创业、教育信息化、教育创业报道等，找教育就上决胜网教育门户网站。",
		ImageUrl:    "https://juesheng.com/s/img/favicon.png",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != ""{

		rssItems := commonParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx,cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
