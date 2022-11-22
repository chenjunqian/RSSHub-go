package dongqiudi

import (
	"context"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetDaily(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	cacheKey := "DONGQIUDI_DAILY"
	if value, err := component.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
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
	if resp := component.GetContent(ctx,apiUrl); resp != "" {
		rssData.Items = commonParser(resp)
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	component.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*6)
	req.Response.WriteXmlExit(rssStr)
}
