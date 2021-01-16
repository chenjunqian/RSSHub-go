package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/mitchina"
)

func MitChinaRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		mitchinaCtl := new(mitchina.Controller)
		group.GET("/flash", mitchinaCtl.GetFlash)
	})
}
