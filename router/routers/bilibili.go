package routers

import (
	"github.com/gogf/gf/net/ghttp"

	"rsshub/app/api/bilibili"
)

func BilibiliRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		bilibiliController := new(bilibili.Controller)
		group.GET("/appversion/:id", bilibiliController.GetAppVersion)
	})
}
