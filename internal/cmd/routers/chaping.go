package routers

import (
	"rsshub/internal/controller/rssapi/chaping"

	"github.com/gogf/gf/v2/net/ghttp"
)

func ChaPingRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/news/game", chaping.Controller.GetNews)
		group.GET("/news/techNews", chaping.Controller.GetNews)
		group.GET("/news/techFun", chaping.Controller.GetNews)
		group.GET("/news/debugTime", chaping.Controller.GetNews)
		group.GET("/news/internetFun", chaping.Controller.GetNews)
		group.GET("/news/car", chaping.Controller.GetNews)
	})
}
