package baijing

import (
	"context"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetTouzi(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	if value, err := service.GetRedis().Do(ctx,"GET", "BAIJING_TOUZI"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.baijingapp.com/index/ajax/getlist/type-10"
	rssData := dao.RSSFeed{
		Title:       "白鲸出海-投融资",
		Link:        "https://www.baijingapp.com",
		Description: "白鲸出海投融资",
		Tag:         []string{"投资"},
		ImageUrl:    "https://www.baijingapp.com/static/css/default/img/favicon.ico",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != "" {
		rssItems := commonHtmlParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	service.GetRedis().Do(ctx,"SET", "BAIJING_TOUZI", rssStr)
	service.GetRedis().Do(ctx,"EXPIRE", "BAIJING_TOUZI", 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
