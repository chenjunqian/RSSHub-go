package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/chaping"
)

func ChaPingRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		chapingCtl := new(chaping.Controller)
		group.GET("/news/game", chapingCtl.GetNews)
		group.GET("/news/techNews", chapingCtl.GetNews)
		group.GET("/news/techFun", chapingCtl.GetNews)
		group.GET("/news/debugTime", chapingCtl.GetNews)
		group.GET("/news/internetFun", chapingCtl.GetNews)
		group.GET("/news/car", chapingCtl.GetNews)
	})
}
