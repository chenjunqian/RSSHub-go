package routers

import (
	"rsshub/internal/controller/rssapi/gcore"

	"github.com/gogf/gf/v2/net/ghttp"
)

func GCoreRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		gcoreCtl := new(gcore.Controller)
		group.GET("news", gcoreCtl.GetIndex)
		group.GET("radios", gcoreCtl.GetIndex)
		group.GET("articles", gcoreCtl.GetIndex)
		group.GET("videos", gcoreCtl.GetIndex)
	})
}
