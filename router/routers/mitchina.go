package routers

import (
	"rsshub/app/api/rssapi/mitchina"

	"github.com/gogf/gf/v2/net/ghttp"
)

func MitChinaRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		mitchinaCtl := new(mitchina.Controller)
		group.GET("/flash", mitchinaCtl.GetFlash)
	})
}
