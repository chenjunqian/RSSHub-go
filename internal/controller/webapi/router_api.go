package webapi

import (
	"rsshub/internal/dao"
	"rsshub/internal/service/feed"
	response "rsshub/internal/service/middleware"
	routerService "rsshub/internal/service/router"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"golang.org/x/net/context"
)

func (ctl *Controller) IndexTpl(req *ghttp.Request) {

	var (
		routerCatInfoList []routerService.CatagoryDirInfo
	)

	routerCatInfoList = routerService.GetRouterCatagoryList()

	req.Response.WriteTpl("home.html", g.Map{
		"name":              "RSS Go",
		"routerCatInfoList": routerCatInfoList,
	})

}

func (ctl *Controller) IndexWithParamTpl(req *ghttp.Request) {

	var (
		routerCatInfoList []routerService.CatagoryDirInfo
		routerDir         string
	)

	routerDir = req.Get("router_dir").String()
	routerCatInfoList = routerService.GetRouterCatagoryList()

	for index, routerCatInfo := range routerCatInfoList {
		if routerDir == routerCatInfo.DirName {
			routerCatInfoList[index].CollapseOpen = true
		} else {
			routerCatInfoList[index].CollapseOpen = false
        }
	}

	req.Response.WriteTpl("home.html", g.Map{
		"name":              "RSS Go",
		"router_dir":        routerDir,
		"routerCatInfoList": routerCatInfoList,
	})

}

func (ctl *Controller) GetAllRssResource(req *ghttp.Request) {

	routerArray := g.Server().GetRoutes()
	routerDataList := make([]RouterInfoData, 0)
	if len(routerArray) > 0 {
		for _, router := range routerArray {
			if strings.Contains(router.Route, ":") || strings.Contains(router.Route, "*") || strings.HasPrefix(router.Route, "/api") {
				continue
			}
			routerInfoData := RouterInfoData{
				Route: router.Route,
				Port:  router.Address,
			}
			routerDataList = append(routerDataList, routerInfoData)
		}
	}

	response.JsonExit(req, 0, "success", routerDataList)
}

func (ctl *Controller) GetAllFeedChannelInfoList(req *ghttp.Request) {

	var feedChannelInfoList []dao.RSSFeed
	var ctx context.Context = context.TODO()

	feedChannelInfoList = feed.GetAllChannelInfoList(ctx)

	response.JsonExit(req, 0, "success", feedChannelInfoList)
}
