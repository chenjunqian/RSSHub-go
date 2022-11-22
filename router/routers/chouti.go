package routers

import (
	"rsshub/app/api/rssapi/chouti"

	"github.com/gogf/gf/v2/net/ghttp"
)

func ChouTiRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/hot", chouti.Controller.GetIndex)
		group.GET("/news", chouti.Controller.GetIndex)
		group.GET("/scoff", chouti.Controller.GetIndex)
		group.GET("/pic", chouti.Controller.GetIndex)
		group.GET("/ask", chouti.Controller.GetIndex)
		group.GET("/tec", chouti.Controller.GetIndex)
		group.GET("/top", chouti.Controller.GetTop)
	})
}
