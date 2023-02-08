package ifeng

import (
	"context"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetTech(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getTechInfoLinks()[linkType]

	cacheKey := "IFENG_TECH_" + linkConfig.ChannelId
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	var apiUrl string
	if linkConfig.ChannelId == "" {
		apiUrl = "https://tech.ifeng.com/"
	} else {
		apiUrl = "https://tech.ifeng.com/" + linkConfig.ChannelId + "/"
	}

	rssData := dao.RSSFeed{
		Title:       "凤凰网|科技 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         []string{"科技"},
		Description: "凤凰新媒体不仅是控股的凤凰卫视传媒集团优质电视内容的网络传播渠道，更整合了众多专业媒体机构生产的内容、用户生成的内容、以及自身生产的专业内容，提供含图文音视频的全方位综合新闻资讯、深度报道、观点评论、财经产品、互动应用、分享社区、在线网页游戏等服务，满足主流人群浏览、表达、交流、分享、娱乐、理财等多元化与个性化的诉求，并反向传输给凤凰卫视的电视平台，形成创新的网台联动组合传播模式，为互联网、移动互联网及视频用户提供丰富的内容与随时随地随身的服务。",
		ImageUrl:    "https://y0.ifengimg.com/index/favicon.ico",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != ""{
		rssData.Items = commonParser(ctx, resp)
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx,cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
