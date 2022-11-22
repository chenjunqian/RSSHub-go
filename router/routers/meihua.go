package routers

import (
	"rsshub/app/api/rssapi/meihua"

	"github.com/gogf/gf/v2/net/ghttp"
)

func MeiHuaRouter(group *ghttp.RouterGroup) {
	group.Group("/index", func(group *ghttp.RouterGroup) {
		meihuaCtl := new(meihua.Controller)
		group.GET("/hot", meihuaCtl.GetIndex)
		group.GET("/latest", meihuaCtl.GetIndex)
	})
}
