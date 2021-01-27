package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/houxu"
)

func HouXuRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		houxuController := new(houxu.Controller)
		group.GET("/index/hot", houxuController.GetIndex)
	})
}
