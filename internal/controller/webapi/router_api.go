package webapi

import (
	"rsshub/internal/dao"
	"rsshub/internal/service/feed"
	response "rsshub/internal/service/middleware"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"golang.org/x/net/context"
)

func (ctl *Controller) IndexTpl(req *ghttp.Request) {

	req.Response.WriteTpl("home.html", g.Map{
		"name": "RSS Go",
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
