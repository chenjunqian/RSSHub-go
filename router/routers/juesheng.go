package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/juesheng"
)

func JueShengRouter(group *ghttp.RouterGroup) {
	group.Group("/index", func(group *ghttp.RouterGroup) {
		group.GET("/news", juesheng.Controller.GetIndex)
		group.GET("/12k", juesheng.Controller.GetIndex)
		group.GET("/e-edu", juesheng.Controller.GetIndex)
		group.GET("/zhijiao", juesheng.Controller.GetIndex)
		group.GET("/xingqu", juesheng.Controller.GetIndex)
		group.GET("/xueqian", juesheng.Controller.GetIndex)
	})
}
