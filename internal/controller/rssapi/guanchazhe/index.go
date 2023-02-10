package guanchazhe

import (
	"context"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]

	cacheKey := "GUANCHAZHE_INDEX_" + linkConfig.ChannelId
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.guancha.cn/" + linkConfig.ChannelId
	rssData := dao.RSSFeed{
		Title:       "观察家 " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         linkConfig.Tags,
		Description: "经济观察网，经济观察报，电子报纸,电子杂志,财经媒体,观察家,eeo",
		ImageUrl:    "https://www.eeo.com.cn/favicon.ico",
	}
	if resp := service.GetContent(ctx, apiUrl); resp != "" {

		var rssItems []dao.RSSItem
		if linkConfig.LinkType == "index" {
			rssItems = indexParser(ctx, resp)
		} else if linkConfig.LinkType == "common" {
			rssItems = commonParser(ctx, resp)
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx, cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
