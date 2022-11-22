package routers

import (
	"rsshub/app/api/rssapi/juesheng"

	"github.com/gogf/gf/v2/net/ghttp"
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
