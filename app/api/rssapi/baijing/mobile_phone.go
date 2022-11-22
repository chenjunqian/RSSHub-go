package baijing

import (
	"context"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetMobilePhone(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	if value, err := component.GetRedis().Do(ctx,"GET", "BAIJING_MOBILE_PHONE"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.baijingapp.com/index/ajax/getlist/type-7"
	rssData := dao.RSSFeed{
		Title:       "白鲸出海-智能手机",
		Link:        "https://www.baijingapp.com",
		Description: "白鲸出海智能手机",
		Tag:         []string{"手机", "移动互联网"},
		ImageUrl:    "https://www.baijingapp.com/static/css/default/img/favicon.ico",
	}
	if resp := component.GetContent(ctx,apiUrl); resp != "" {
		rssItems := commonHtmlParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx,"SET", "BAIJING_MOBILE_PHONE", rssStr)
	component.GetRedis().Do(ctx,"EXPIRE", "BAIJING_MOBILE_PHONE", 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
