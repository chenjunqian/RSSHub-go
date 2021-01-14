package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/cnbeta"
)

func CNBetaRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		cnbetaCtl := new(cnbeta.Controller)
		group.GET("/", cnbetaCtl.GetRSSSource)
	})
}
