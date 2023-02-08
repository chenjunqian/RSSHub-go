package dx2025

import (
	"context"
	"fmt"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetCategoryNews(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	routeArray := strings.Split(req.Router.Uri, "/")
	var linkConfig IndustryInfoRouteConfig
	var categoryType string
	if len(routeArray) > 2 {
		linkType := routeArray[len(routeArray)-1]
		categoryType = routeArray[len(routeArray)-2]
		linkConfig = getIndustryInfoLinks()[linkType]
	}

	cacheKey := fmt.Sprintf("DX2025_%s_%s", categoryType, linkConfig.ChannelId)
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	var dxRouteType string
	var apiUrl string
	var feedTitle string
	if categoryType == "report" {
		dxRouteType = "industry-reports"
		apiUrl = fmt.Sprintf("https://www.dx2025.com/archives/category/%s/%s-%s", dxRouteType, linkConfig.ChannelId, dxRouteType)
		feedTitle = "东西智库 - 行业报告 - " + linkConfig.Title
	} else if categoryType == "observation" {
		dxRouteType = "industry-observation"
		apiUrl = fmt.Sprintf("https://www.dx2025.com/archives/category/%s/%s", dxRouteType, linkConfig.ChannelId)
		feedTitle = "东西智库 - 产业观察 - " + linkConfig.Title
	}
	rssData := dao.RSSFeed{
		Title:       feedTitle,
		Link:        apiUrl,
		Description: "东西智库专注中国制造业高质量发展",
		Tag:         []string{"其他"},
		ImageUrl:    "https://www.dx2025.com/wp-content/uploads/2020/04/cropped-east_west_think_tank_800x800-32x32.png",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != "" {
		rssData.Items = commonParser(ctx, resp)
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx,cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
