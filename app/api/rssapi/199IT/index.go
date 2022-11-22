package _199IT

import (
	"context"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) Get199ITIndex(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	if value, err := component.GetRedis().Do(ctx,"GET", "199IT_INDEX"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "http://www.199it.com/newly"
	rssData := dao.RSSFeed{
		Title:       "199it",
		Link:        apiUrl,
		ImageUrl:    "https://www.199it.com/favicon.ico",
		Tag:         []string{"互联网", "IT", "科技"},
		Description: "互联网数据资讯网-199IT",
	}
	if resp := component.GetContent(ctx,apiUrl); resp != "" {
		rssItems := parseArticle(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx,"SET", "199IT_INDEX", rssStr)
	component.GetRedis().Do(ctx,"EXPIRE", "199IT_INDEX", 60*10)
	req.Response.WriteXmlExit(rssStr)
}
