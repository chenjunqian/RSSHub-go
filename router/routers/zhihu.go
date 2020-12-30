package routers

import (
	"rsshub/app/api/zhihu"

	"github.com/gogf/gf/net/ghttp"
)

func ZhihubRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		zhihuController := new(zhihu.Controller)
		group.GET("/activities/:id", zhihuController.GetActivities)
		group.GET("/answers/:id", zhihuController.GetAnswers)
		group.GET("/collections/:id", zhihuController.GetCollections)
		group.GET("/zhuanlan/:id", zhihuController.GetZhuanlan)
	})
}
