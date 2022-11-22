package routers

import (
	"rsshub/app/api/rssapi/ccg"

	"github.com/gogf/gf/v2/net/ghttp"
)

func CCGRouter(group *ghttp.RouterGroup) {
	group.Group("/index", func(group *ghttp.RouterGroup) {
		group.GET("/news", ccg.Controller.GetIndex)
		group.GET("/media", ccg.Controller.GetIndex)
	})
}
