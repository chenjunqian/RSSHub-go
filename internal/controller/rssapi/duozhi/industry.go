package duozhi

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetIndustryNews(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getIndustryNewsLinks()[linkType]

	cacheKey := "DUOZHI_INDUSTRY_" + linkConfig.ChannelId
	if value, err := service.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "http://www.duozhi.com/industry/" + linkConfig.ChannelId
	rssData := dao.RSSFeed{
		Title:       "多知 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         linkConfig.Tags,
		Description: "多知网 - 独立商业视角 新锐教育观察",
		ImageUrl:    "http://www.duozhi.com/favicon.ico",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != "" {
		rssData.Items = commonParser(ctx, resp)
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	service.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	service.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*6)
	req.Response.WriteXmlExit(rssStr)
}
