package dx2025

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
	"strings"
)

func (ctl *Controller) GetCategoryNews(req *ghttp.Request) {
	routeArray := strings.Split(req.Router.Uri, "/")
	var linkConfig IndustryInfoRouteConfig
	var categoryType string
	if len(routeArray) > 2 {
		linkType := routeArray[len(routeArray)-1]
		categoryType = routeArray[len(routeArray)-2]
		linkConfig = getIndustryInfoLinks()[linkType]
	}

	cacheKey := fmt.Sprintf("DX2025_%s_%s", categoryType, linkConfig.ChannelId)
	//if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
	//	if value.String() != "" {
	//		_ = req.Response.WriteXmlExit(value.String())
	//	}
	//}
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
		ImageUrl:    "https://www.dx2025.com/wp-content/uploads/2020/04/cropped-east_west_think_tank_800x800-32x32.png",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssData.Items = commonParser(resp.ReadAllString())
	}

	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
