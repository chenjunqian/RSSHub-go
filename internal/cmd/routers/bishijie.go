package routers

import (
	"rsshub/internal/controller/rssapi/bishijie"

	"github.com/gogf/gf/v2/net/ghttp"
)

func BiShiJieRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		biShiJieCtl := new(bishijie.Controller)
		group.GET("/flash", biShiJieCtl.GetFlash)
		group.GET("/shendu", biShiJieCtl.GetShenDu)
	})
}
