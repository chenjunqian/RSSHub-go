package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/web"
)

func WebRouter(group *ghttp.RouterGroup) {
	group.Group("/v1", func(group *ghttp.RouterGroup) {
		webController := new(web.Controller)
		group.GET("/routers", webController.GetAllRssResource)
	})
}
