package whalegogo

import (
	"context"
	"fmt"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetIPortal(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]

	cacheKey := "WHALE_GOGO_PORTAL_" + linkConfig.ChannelId
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := fmt.Sprintf("https://api.whalegogo.com/v1/articles/search?tagid=&type_id=%s&sort=-push_at&page=1&per-page=12", linkConfig.ChannelId)
	rssData := dao.RSSFeed{
		Title:       "鲸跃汽车 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         []string{"汽车"},
		Description: "我们是一帮在传统汽车门户、汽车杂志战斗过数年的老司机，关于车与理想生活的种种，我们有很多思考，现在用一个全新的网站和 APP，将车型快讯、试驾体验、行业深度、生活方式等原创内容，用简洁的设计与排版呈现在你的面前。",
		ImageUrl:    "https://api.whalegogo.com/favicon.ico",
	}
	if resp := service.GetContent(ctx, apiUrl); resp != "" {

		rssItems := portalParser(resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx, cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
