package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/webapi"
)

func WebRouter(group *ghttp.RouterGroup) {
	group.Group("/v1", func(group *ghttp.RouterGroup) {
		//group.Middleware(middleware.AuthToken)
		webController := new(webapi.Controller)
		group.Group("", func(group *ghttp.RouterGroup) {
			group.GET("/routers", webController.GetAllRssResource)
		})
	})
}
