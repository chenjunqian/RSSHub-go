package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/sspai"
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
