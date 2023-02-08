package baijing

import (
	"context"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetActivity(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	if value, err := cache.GetCache(ctx, "BAIJING_ACTIVITY"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.baijingapp.com/index/ajax/getlist/type-5"
	rssData := dao.RSSFeed{
		Title:       "白鲸出海-活动",
		Link:        "https://www.baijingapp.com",
		Description: "白鲸出海活动",
		Tag:         []string{"互联网", "开发者", "科技", "社区"},
		ImageUrl:    "https://www.baijingapp.com/static/css/default/img/favicon.ico",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != "" {
		rssItems := commonHtmlParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx,"BAIJING_ACTIVITY", rssStr)
	req.Response.WriteXmlExit(rssStr)
}
