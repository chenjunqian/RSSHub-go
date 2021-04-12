package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/bishijie"
)

func BiShiJieRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		biShiJieCtl := new(bishijie.Controller)
		group.GET("/flash", biShiJieCtl.GetFlash)
		group.GET("/shendu", biShiJieCtl.GetShenDu)
	})
}
