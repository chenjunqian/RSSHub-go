package niaogenote

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetCat(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]

	cacheKey := "NGBJ_CAT_" + linkConfig.ChannelId
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.niaogebiji.com/cat/" + linkConfig.ChannelId
	rssData := dao.RSSFeed{
		Title:       "鸟哥笔记 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         []string{"运营"},
		Description: "鸟哥笔记新媒体运营栏目，提供微信、微博、贴吧等新兴媒体平台运营策略，研究如何通过现代化互联网手段进行产品宣传、推广、产品营销，如何策划品牌相关的优质、高度传播性的内容和线上活动，如何向客户广泛或者精准推送消息，如何充分利用粉丝经济，达到相应营销目的。",
		ImageUrl:    "https://www.niaogebiji.com/favicon.ico",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != ""{

		rssItems := catParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx,cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
