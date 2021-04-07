package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/web"
)

func WebRouter(group *ghttp.RouterGroup) {
	group.Group("/v1", func(group *ghttp.RouterGroup) {
		//group.Middleware(middleware.AuthToken)
		webController := new(web.Controller)
		group.Group("", func(group *ghttp.RouterGroup) {
			group.GET("/routers", webController.GetAllRssResource)
			group.GET("/feed_tag", webController.GetFeedTag)
		})
	})
}
