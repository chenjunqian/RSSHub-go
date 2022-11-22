package guanchajia

import (
	"context"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]

	cacheKey := "GUANCHAJIA_INDEX_" + linkConfig.ChannelId
	if value, err := component.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "http://www.eeo.com.cn/" + linkConfig.ChannelId
	rssData := dao.RSSFeed{
		Title:       "观察家 " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         linkConfig.Tags,
		Description: "经济观察网，经济观察报，电子报纸,电子杂志,财经媒体,观察家,eeo",
		ImageUrl:    "https://www.eeo.com.cn/favicon.ico",
	}
	if resp := component.GetContent(ctx,apiUrl); resp != ""{

		rssItems := commonParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	component.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
