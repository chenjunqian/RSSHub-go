package routers

import (
	"rsshub/app/api/zhihu"

	"github.com/gogf/gf/net/ghttp"
)

func ZhihubRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		zhihuController := new(zhihu.Controller)
		group.GET("/activities", zhihuController.GetActivities)
		group.GET("/answers", zhihuController.GetAnswers)
	})
}
