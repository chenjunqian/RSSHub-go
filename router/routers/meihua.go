package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/meihua"
)

func MeiHuaRouter(group *ghttp.RouterGroup) {
	group.Group("/index", func(group *ghttp.RouterGroup) {
		meihuaCtl := new(meihua.Controller)
		group.GET("/hot", meihuaCtl.GetIndex)
		group.GET("/latest", meihuaCtl.GetIndex)
	})
}
