package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/ifan"
)

func IFanRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		ifanCtl := new(ifan.Controller)
		group.GET("/daily", ifanCtl.GetFlash)
	})
}
