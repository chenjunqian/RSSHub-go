package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/juesheng"
)

func JueShengRouter(group *ghttp.RouterGroup) {
	group.Group("/index", func(group *ghttp.RouterGroup) {
		jueshengCtl := new(juesheng.Controller)
		group.GET("/news", jueshengCtl.GetIndex)
		group.GET("/12k", jueshengCtl.GetIndex)
		group.GET("/e-edu", jueshengCtl.GetIndex)
		group.GET("/zhijiao", jueshengCtl.GetIndex)
		group.GET("/xingqu", jueshengCtl.GetIndex)
		group.GET("/xueqian", jueshengCtl.GetIndex)
	})
}
