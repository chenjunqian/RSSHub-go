package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/fulinian"
)

func FuLiNianRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		fulinianCtl := new(fulinian.Controller)
		group.GET("index", fulinianCtl.GetIndex)
		group.GET("technical-course", fulinianCtl.GetIndex)
		group.GET("learning", fulinianCtl.GetIndex)
		group.GET("chuangye", fulinianCtl.GetIndex)
		group.GET("fulinian", fulinianCtl.GetIndex)
		group.GET("network-resource", fulinianCtl.GetIndex)
		group.GET("quality-software", fulinianCtl.GetIndex)
	})
}
