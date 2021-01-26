package guanchazhe

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

	cacheKey := "GUANCHAZHE_INDEX_" + linkConfig.ChannelId
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.guancha.cn/" + linkConfig.ChannelId
	rssData := dao.RSSFeed{
		Title:       "观察家 " + linkConfig.Title,
		Link:        apiUrl,
		Description: "经济观察网，经济观察报，电子报纸,电子杂志,财经媒体,观察家,eeo",
		ImageUrl:    "http://www.eeo.com.cn/favicon.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		var rssItems []dao.RSSItem
		if linkConfig.LinkType == "index" {
			rssItems = indexParser(resp.ReadAllString())
		} else if linkConfig.LinkType == "common" {
			rssItems = commonParser(resp.ReadAllString())
		}
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
