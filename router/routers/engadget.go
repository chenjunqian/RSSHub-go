package routers

import (
	"rsshub/app/api/rssapi/engadget"

	"github.com/gogf/gf/v2/net/ghttp"
)

func EngadgetRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		engadgetCtl := new(engadget.Controller)
		group.GET("/index", engadgetCtl.GetIndexRSS)
	})
}
