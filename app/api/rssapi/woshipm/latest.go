package woshipm

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {

	cacheKey := "WOSHIPM_LATEST"
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "http://www.woshipm.com/__api/v1/stream-list"
	rssData := dao.RSSFeed{
		Title:       "人人都是产品经理 - 最新文章",
		Link:        apiUrl,
		Tag:         []string{"互联网"},
		Description: "人人都是产品经理（woshipm.com）是以产品经理、运营为核心的学习、交流、分享平台，集媒体、培训、社群为一体，全方位服务产品人和运营人，成立9年举办在线讲座500+期，线下分享会300+场，产品经理大会、运营大会20+场，覆盖北上广深杭成都等15个城市，在行业有较高的影响力和知名度。平台聚集了众多BAT美团京东滴滴360小米网易等知名互联网公司产品总监和运营总监，他们在这里与你一起成长。",
		ImageUrl:    "https://image.woshipm.com/favicon.ico",
	}

	if resp, err := component.GetHttpClient().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := commonParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
