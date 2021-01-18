package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/chouti"
)

func ChouTiRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		chouTiCtl := new(chouti.Controller)
		group.GET("/hot", chouTiCtl.GetIndex)
		group.GET("/news", chouTiCtl.GetIndex)
		group.GET("/scoff", chouTiCtl.GetIndex)
		group.GET("/pic", chouTiCtl.GetIndex)
		group.GET("/ask", chouTiCtl.GetIndex)
		group.GET("/tec", chouTiCtl.GetIndex)
		group.GET("/top", chouTiCtl.GetTop)
	})
}
