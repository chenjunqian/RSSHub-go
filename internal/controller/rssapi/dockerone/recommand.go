package dockerone

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetRecommand(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	cacheKey := "DOCKERONE_RECOMMAND"
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "http://weekly.dockone.io/is_recommend-1"
	rssData := dao.RSSFeed{
		Title:       "Dockone",
		Link:        apiUrl,
		Description: "DockOne.io,为技术人员提供最专业的Cloud Native交流平台。",
		ImageUrl:    "http://weekly.dockone.io/static/css/default/img/favicon.ico",
	}
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		rssItems := parseRecommand(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx, cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
