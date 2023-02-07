package routers

import (
	"rsshub/internal/controller/rssapi/oschina"

	"github.com/gogf/gf/v2/net/ghttp"
)

func OSChinaRouter(group *ghttp.RouterGroup) {
	group.Group("/news", func(group *ghttp.RouterGroup) {
		group.GET("/latest", oschina.Controller.GetLatestNews)
	})
}