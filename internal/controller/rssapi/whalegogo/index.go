package whalegogo

import (
	"context"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	cacheKey := "WHALE_GOGO_INDEX"
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://api.whalegogo.com/v1/app/index"
	rssData := dao.RSSFeed{
		Title:       "鲸跃汽车 - 最新",
		Tag:         []string{"汽车"},
		Link:        "https://m.whalegogo.com/index",
		Description: "我们是一帮在传统汽车门户、汽车杂志战斗过数年的老司机，关于车与理想生活的种种，我们有很多思考，现在用一个全新的网站和 APP，将车型快讯、试驾体验、行业深度、生活方式等原创内容，用简洁的设计与排版呈现在你的面前。",
		ImageUrl:    "https://api.whalegogo.com/favicon.ico",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != ""{

		rssItems := indexParser(ctx,resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx,cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
