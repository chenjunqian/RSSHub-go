package _36kr

import (
	"context"
	"rsshub/internal/service/feed"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) Get36krNews(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getNewsLinks()[linkType]

	if value, err := service.GetRedis().Do(ctx,"GET", "36KR_NEWS_"+linkConfig.Link); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://36kr.com" + linkConfig.Link
	rssData := dao.RSSFeed{
		Title:    "36氪资讯 - " + linkConfig.Title,
		Link:     apiUrl,
		Tag:      linkConfig.Tags,
		ImageUrl: "https://static.36krcdn.com/36kr-web/static/ic_default_100_56@2x.ec858a2a.png",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != "" {
		rssItems := parseNews(ctx, resp)
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	service.GetRedis().Do(ctx,"SET", "36KR_NEWS_"+linkConfig.Link, rssStr)
	service.GetRedis().Do(ctx,"EXPIRE", "36KR_NEWS_"+linkConfig.Link, 60*60*3)
	req.Response.WriteXmlExit(rssStr)
}
