package sspai

import (
	"fmt"
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
	cacheKey := "SSPAI_INDEX_" + linkConfig.ChannelId
	//if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
	//	if value.String() != "" {
	//		_ = req.Response.WriteXmlExit(value.String())
	//	}
	//}
	var apiUrl string
	if linkConfig.ChannelId == "recommend" {
		apiUrl = "https://sspai.com/api/v1/article/index/page/get?limit=10&offset=0&created_at=0"
	} else {
		apiUrl = fmt.Sprintf("https://sspai.com/api/v1/article/tag/page/get?limit=10&offset=0&tag=%s", linkConfig.Title)
	}
	rssData := dao.RSSFeed{
		Title:       "少数派 - " + linkConfig.Title,
		Link:        apiUrl,
		Description: "少数派致力于更好地运用数字产品或科学方法，帮助用户提升工作效率和生活品质",
		ImageUrl:    "https://cdn.sspai.com/sspai/assets/img/favicon/icon.ico",
	}

	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := commonParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
