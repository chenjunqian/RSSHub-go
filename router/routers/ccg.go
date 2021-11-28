package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/ccg"
)

func CCGRouter(group *ghttp.RouterGroup) {
	group.Group("/index", func(group *ghttp.RouterGroup) {
		group.GET("/news", ccg.Controller.GetIndex)
		group.GET("/media", ccg.Controller.GetIndex)
	})
}
