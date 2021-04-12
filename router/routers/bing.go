package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/bing"
)

func BingRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		bingController := new(bing.Controller)
		group.GET("/daily-image/", bingController.GetDailyImage)
	})
}
