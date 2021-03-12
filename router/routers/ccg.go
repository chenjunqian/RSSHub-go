package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/ccg"
)

func CCGRouter(group *ghttp.RouterGroup) {
	group.Group("/index", func(group *ghttp.RouterGroup) {
		ccgCtl := new(ccg.Controller)
		group.GET("/news", ccgCtl.GetIndex)
		group.GET("/media", ccgCtl.GetIndex)
	})
}
