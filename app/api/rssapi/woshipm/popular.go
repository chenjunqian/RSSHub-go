package woshipm

import (
	"context"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetPopular(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	cacheKey := "WOSHIPM_POPULAR"
	if value, err := component.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "http://www.woshipm.com/__api/v1/browser/popular"
	rssData := dao.RSSFeed{
		Title:       "人人都是产品经理 - 热门文章",
		Link:        apiUrl,
		Tag:         []string{"科技", "媒体", "运营"},
		Description: "人人都是产品经理（woshipm.com）是以产品经理、运营为核心的学习、交流、分享平台，集媒体、培训、社群为一体，全方位服务产品人和运营人，成立9年举办在线讲座500+期，线下分享会300+场，产品经理大会、运营大会20+场，覆盖北上广深杭成都等15个城市，在行业有较高的影响力和知名度。平台聚集了众多BAT美团京东滴滴360小米网易等知名互联网公司产品总监和运营总监，他们在这里与你一起成长。",
		ImageUrl:    "https://image.woshipm.com/favicon.ico",
	}

	if resp := component.GetContent(ctx,apiUrl); resp != ""{
		rssItems := commonParser(ctx, resp)
		rssData.Items = rssItems

	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	component.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
