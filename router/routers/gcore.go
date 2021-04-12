package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/gcore"
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
