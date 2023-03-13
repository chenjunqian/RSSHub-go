package routers

import (
	"rsshub/internal/controller/rssapi/sspai"

	"github.com/gogf/gf/v2/net/ghttp"
)

func SSPaiRouter(group *ghttp.RouterGroup) {
	group.Group("/index", func(group *ghttp.RouterGroup) {
		sspaiCtl := new(sspai.Controller)
		group.GET("/recommend", sspaiCtl.GetIndex)
		group.GET("/hot", sspaiCtl.GetIndex)
		group.GET("/app_recommend", sspaiCtl.GetIndex)
		group.GET("/skill", sspaiCtl.GetIndex)
		group.GET("/lifestyle", sspaiCtl.GetIndex)
		group.GET("/podcast", sspaiCtl.GetIndex)
	})
}
