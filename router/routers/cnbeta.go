package routers

import (
	"rsshub/app/api/rssapi/cnbeta"

	"github.com/gogf/gf/v2/net/ghttp"
)

func CNBetaRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		cnbetaCtl := new(cnbeta.Controller)
		group.GET("/", cnbetaCtl.GetRSSSource)
	})
}
