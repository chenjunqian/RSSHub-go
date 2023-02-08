package medsci

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]

	cacheKey := "MEDSCI_RECOMMEDN_INDEX_" + linkConfig.ChannelId
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.medsci.cn/"
	rssData := dao.RSSFeed{
		Title:       "梅斯 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         []string{"医疗"},
		Description: "MedSci(梅斯医学)致力于医疗质量的改善，从事临床研究服务、数据管理、医学统计、临床培训、继续教育等支持，促进临床医生职业发展和医疗智慧化。",
		ImageUrl:    "https://cache1.medsci.cn/images/favicon.ico",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != ""{
		rssItems := commonParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx,cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
