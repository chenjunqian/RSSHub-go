package dockerone

import (
	"context"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetRecommand(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	cacheKey := "DOCKERONE_RECOMMAND"
	if value, err := component.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
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
	if resp := component.GetContent(ctx,apiUrl); resp != "" {
		rssItems := parseRecommand(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	component.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}