package juesheng

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
	"strings"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]
	cacheKey := "JUESHENG_INDEX_" + linkConfig.ChannelUrl
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := linkConfig.ChannelUrl
	rssData := dao.RSSFeed{
		Title:       "决胜网 - " + linkConfig.Title,
		Link:        apiUrl,
		Description: "决胜网是教育产业门户网站提供：教育门户新闻资讯、互联网+教育、在线教育、兴趣教育、在线职业教育、教育创业、教育信息化、教育创业报道等，找教育就上决胜网教育门户网站。",
		ImageUrl:    "https://juesheng.com/s/img/favicon.png",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := commonParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
