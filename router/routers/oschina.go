package routers

import (
	"rsshub/app/api/rssapi/oschina"

	"github.com/gogf/gf/net/ghttp"
)

func OSChinaRouter(group *ghttp.RouterGroup) {
	group.Group("/news", func(group *ghttp.RouterGroup) {
		group.GET("/latest", oschina.Controller.GetLatestNews)
	})
}