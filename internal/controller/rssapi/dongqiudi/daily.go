package dongqiudi

import (
	"context"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetDaily(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	cacheKey := "DONGQIUDI_DAILY"
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.dongqiudi.com/special/48"
	rssData := dao.RSSFeed{
		Title:       "懂球帝 - 早报",
		Link:        apiUrl,
		Tag:         []string{"体育"},
		Description: "早报 — 专题|专业权威的足球网站|懂球帝",
		ImageUrl:    "https://static1.dongqiudi.com/web-new/web/images/fav.ico",
	}
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		rssData.Items = commonParser(resp)
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx, cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
