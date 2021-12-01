package gcore

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]

	cacheKey := "GCORE_" + linkConfig.ChannelId
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.gcores.com/" + linkConfig.ChannelId
	rssData := dao.RSSFeed{
		Title:       "机核 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         linkConfig.Tags,
		Description: "机核从2010年开始一直致力于分享游戏玩家的生活，以及深入探讨游戏相关的文化。我们开发原创的电台以及视频节目，一直在不断寻找民间高质量的内容创作者。",
		ImageUrl:    "https://www.gcores.com/favicon.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := commonParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
