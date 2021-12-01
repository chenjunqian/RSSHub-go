package niaogenote

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"
)

func (ctl *Controller) GetCat(req *ghttp.Request) {

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]

	cacheKey := "NGBJ_CAT_" + linkConfig.ChannelId
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.niaogebiji.com/cat/" + linkConfig.ChannelId
	rssData := dao.RSSFeed{
		Title:       "鸟哥笔记 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         []string{"运营"},
		Description: "鸟哥笔记新媒体运营栏目，提供微信、微博、贴吧等新兴媒体平台运营策略，研究如何通过现代化互联网手段进行产品宣传、推广、产品营销，如何策划品牌相关的优质、高度传播性的内容和线上活动，如何向客户广泛或者精准推送消息，如何充分利用粉丝经济，达到相应营销目的。",
		ImageUrl:    "https://www.niaogebiji.com/favicon.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := catParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
