package routers

import (
	"rsshub/internal/controller/rssapi/houxu"

	"github.com/gogf/gf/v2/net/ghttp"
)

func HouXuRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		houxuController := new(houxu.Controller)
		group.GET("/index/hot", houxuController.GetIndex)
	})
}
