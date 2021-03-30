package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/web"
	"rsshub/app/service/middleware"
)

func WebRouter(group *ghttp.RouterGroup) {
	group.Group("/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.AuthToken)
		webController := new(web.Controller)
		group.GET("/routers", webController.GetAllRssResource)
	})
}
