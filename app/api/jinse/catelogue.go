package jinse

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
	"strings"
)

func (ctl *Controller) GetCatalogue(req *ghttp.Request) {

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]

	cacheKey := "JINSE_CATELOGUE_" + linkConfig.ChannelId
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := fmt.Sprintf("https://api.jinse.com/v6/www/information/list?catelogue_key=%s&limit=23&flag=down", linkConfig.ChannelId)
	rssData := dao.RSSFeed{
		Title:       "金色财经 - " + linkConfig.Title,
		Link:        "https://www.jinse.com/",
		Description: "金色财经是集行业新闻、资讯、行情、数据等一站式区块链产业服务平台，我们追求及时、全面、专业、准确的资讯与数据，致力于为区块链创业者以及数字货币投资者提供最好的产品和服务。",
		ImageUrl:    "https://www.gcores.com/favicon.ico?",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := catalogueParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
