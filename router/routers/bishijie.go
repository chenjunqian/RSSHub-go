package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/bishijie"
)

func BiShiJieRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		biShiJieCtl := new(bishijie.Controller)
		group.GET("/flash", biShiJieCtl.GetFlash)
	})
}
