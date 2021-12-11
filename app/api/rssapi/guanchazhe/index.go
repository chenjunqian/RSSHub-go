package guanchazhe

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
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
		Tag:         linkConfig.Tags,
		Description: "经济观察网，经济观察报，电子报纸,电子杂志,财经媒体,观察家,eeo",
		ImageUrl:    "https://www.eeo.com.cn/favicon.ico",
	}
	if resp, err := component.GetHttpClient().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		defer func(resp *ghttp.ClientResponse) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(err)
			}
		}(resp)
		var rssItems []dao.RSSItem
		if linkConfig.LinkType == "index" {
			rssItems = indexParser(resp.ReadAllString())
		} else if linkConfig.LinkType == "common" {
			rssItems = commonParser(resp.ReadAllString())
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
