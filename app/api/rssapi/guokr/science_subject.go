package guokr

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetScienceSubject(req *ghttp.Request) {

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]

	cacheKey := "GUOKR_SCIENCE_SUBJECT" + linkConfig.ChannelId
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.guokr.com/beta/proxy/science_api/articles?retrieve_type=by_subject&page=1&subject_key=" + linkConfig.ChannelId
	rssData := dao.RSSFeed{
		Title:       "果壳 - 学科 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         linkConfig.Tags,
		Description: "果壳网是一个泛科技主题网站，提供负责任、有智趣、贴近生活的内容，你可以在这里阅读、分享、交流、提问。果壳网致力于让科技兴趣成为人们文化生活和娱乐生活的重要元素。",
		ImageUrl:    "https://www.guokr.com/favicon.ico",
	}
	if resp := component.GetContent(apiUrl); resp != ""{

		rssItems := commonParser(resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
