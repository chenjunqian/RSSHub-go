package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/pintu"
)

func PintuRouter(group *ghttp.RouterGroup) {
	group.Group("/index", func(group *ghttp.RouterGroup) {
		pintuCtl := new(pintu.Controller)
		group.GET("/recommend", pintuCtl.GetIndex)
		group.GET("/sell", pintuCtl.GetIndex)
		group.GET("/tech", pintuCtl.GetIndex)
		group.GET("/entertainment", pintuCtl.GetIndex)
		group.GET("/edu", pintuCtl.GetIndex)
		group.GET("/health", pintuCtl.GetIndex)
		group.GET("/consume", pintuCtl.GetIndex)
		group.GET("/startup", pintuCtl.GetIndex)
	})
}
