package jinse

import (
	"context"
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetCatalogue(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]

	cacheKey := "JINSE_CATELOGUE_" + linkConfig.ChannelId
	if value, err := component.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := fmt.Sprintf("https://api.jinse.com/v6/www/information/list?catelogue_key=%s&limit=23&flag=down", linkConfig.ChannelId)
	rssData := dao.RSSFeed{
		Title:       "金色财经 - " + linkConfig.Title,
		Link:        "https://www.jinse.com/",
		Description: "金色财经是集行业新闻、资讯、行情、数据等一站式区块链产业服务平台，我们追求及时、全面、专业、准确的资讯与数据，致力于为区块链创业者以及数字货币投资者提供最好的产品和服务。",
		ImageUrl:    "https://www.jinse.com/favicon.ico",
	}
	if resp := component.GetContent(ctx,apiUrl); resp != ""{

		rssItems := catalogueParser(ctx,resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	component.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
