package dongqiudi

import (
	"context"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetSpecial(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getNewsLinks()[linkType]

	cacheKey := "DONGQIUDI_SPECIAL_" + linkConfig.ChannelId
	if value, err := component.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.dongqiudi.com/special/" + linkConfig.ChannelId
	rssData := dao.RSSFeed{
		Title:       "懂球帝专题 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         linkConfig.Tags,
		Description: "懂球帝专题|专业权威的足球网站",
		ImageUrl:    "https://static1.dongqiudi.com/web-new/web/images/fav.ico",
	}
	if resp := component.GetContent(ctx,apiUrl); resp != "" {
		rssData.Items = commonParser(resp)
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	component.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*6)
	req.Response.WriteXmlExit(rssStr)
}
