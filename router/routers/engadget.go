package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/engadget"
)

func EngadgetRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		engadgetCtl := new(engadget.Controller)
		group.GET("/index", engadgetCtl.GetIndexRSS)
	})
}
