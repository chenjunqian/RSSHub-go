package baijing

import (
	"context"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetZhuanlan(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	if value, err := service.GetRedis().Do(ctx,"GET", "BAIJING_ZHUANLAN"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.baijingapp.com/index/ajax/getlist/type-4"
	rssData := dao.RSSFeed{
		Title:       "白鲸出海-专栏",
		Link:        "https://www.baijingapp.com",
		Description: "白鲸出海专栏",
		Tag:         []string{"新闻"},
		ImageUrl:    "https://www.baijingapp.com/static/css/default/img/favicon.ico",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != "" {
		rssItems := commonHtmlParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	service.GetRedis().Do(ctx,"SET", "BAIJING_ZHUANLAN", rssStr)
	service.GetRedis().Do(ctx,"EXPIRE", "BAIJING_ZHUANLAN", 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
