package yanxishe

import (
	"context"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetIndex(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	cacheKey := "YANXISHE_ARTICLE"
	if value, err := component.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://api.yanxishe.com/"
	rssData := dao.RSSFeed{
		Title:       "研习社",
		Link:        apiUrl,
		Description: "雷峰网成立于2011年,秉承“关注智能与未来”的宗旨,持续对全球前沿技术趋势与产品动态进行深入调研与解读,是国内具有代表性的实力型科技新媒体与信息服务平台.",
		ImageUrl:    "https://api.yanxishe.com/favicon.ico",
	}
	if resp := component.GetContent(ctx,apiUrl); resp != "" {
		rssItems := parseIndex(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	component.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
