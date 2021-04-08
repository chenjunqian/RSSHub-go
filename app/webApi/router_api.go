package webApi

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	response "rsshub/middleware"
	"strings"
)

func (ctl *Controller) GetAllRssResource(req *ghttp.Request) {

	routerArray := g.Server().GetRouterArray()
	routerDataList := make([]RouterInfoData, 0)
	if len(routerArray) > 0 {
		for _, router := range routerArray {
			if strings.Contains(router.Route, ":") {
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
