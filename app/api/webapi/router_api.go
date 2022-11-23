package webapi

import (
	response "rsshub/middleware"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

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
