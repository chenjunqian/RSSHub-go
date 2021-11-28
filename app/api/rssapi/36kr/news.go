package _36kr

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
	"strings"
)

func (ctl *controller) Get36krNews(req *ghttp.Request) {
	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getNewsLinks()[linkType]

	if value, err := g.Redis().DoVar("GET", "36KR_NEWS_"+linkConfig.Link); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://36kr.com" + linkConfig.Link
	rssData := dao.RSSFeed{
		Title:    "36氪资讯 - " + linkConfig.Title,
		Link:     apiUrl,
		Tag:      linkConfig.Tags,
		ImageUrl: "https://static.36krcdn.com/36kr-web/static/ic_default_100_56@2x.ec858a2a.png",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := parseNews(resp.ReadAllString())
		rssData.Items = rssItems
	}
	rssStr := lib.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "36KR_NEWS_"+linkConfig.Link, rssStr)
	g.Redis().DoVar("EXPIRE", "36KR_NEWS_"+linkConfig.Link, 60*60*3)
	_ = req.Response.WriteXmlExit(rssStr)
}
