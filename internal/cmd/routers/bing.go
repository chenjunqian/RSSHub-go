package routers

import (
	"rsshub/internal/controller/rssapi/bing"

	"github.com/gogf/gf/v2/net/ghttp"
)

func BingRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		bingController := new(bing.Controller)
		group.GET("/daily-image/", bingController.GetDailyImage)
	})
}
