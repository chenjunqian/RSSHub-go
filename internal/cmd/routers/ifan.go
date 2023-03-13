package routers

import (
	"rsshub/internal/controller/rssapi/ifan"

	"github.com/gogf/gf/v2/net/ghttp"
)

func IFanRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		ifanCtl := new(ifan.Controller)
		group.GET("/daily", ifanCtl.GetFlash)
	})
}
