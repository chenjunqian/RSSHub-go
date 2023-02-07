package baijing

import (
	"context"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetEBusiness(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	if value, err := service.GetRedis().Do(ctx,"GET", "BAIJING_EBUSINESS"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.baijingapp.com/index/ajax/getlist/type-5"
	rssData := dao.RSSFeed{
		Title:       "白鲸出海-电商",
		Link:        "https://www.baijingapp.com",
		Description: "白鲸出海电商",
		Tag:         []string{"电商", "数据"},
		ImageUrl:    "https://www.baijingapp.com/static/css/default/img/favicon.ico",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != "" {
		rssItems := commonHtmlParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	service.GetRedis().Do(ctx,"SET", "BAIJING_EBUSINESS", rssStr)
	service.GetRedis().Do(ctx,"EXPIRE", "BAIJING_EBUSINESS", 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
