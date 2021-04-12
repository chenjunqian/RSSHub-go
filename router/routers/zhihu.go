package routers

import (
	"rsshub/app/api/rssapi/zhihu"

	"github.com/gogf/gf/net/ghttp"
)

func ZhihubRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		zhihuController := new(zhihu.Controller)
		group.GET("/activities/:id", zhihuController.GetActivities)
		group.GET("/answers/:id", zhihuController.GetAnswers)
		group.GET("/collections/:id", zhihuController.GetCollections)
		group.GET("/zhuanlan/:id", zhihuController.GetZhuanlan)
		group.GET("/daily", zhihuController.GetDaily)
		group.GET("/daily/section/:id", zhihuController.GetZhihuDailySection)
		group.GET("/daily/hotlist", zhihuController.GetZhihuHostList)
		group.GET("/pin/daily", zhihuController.GetZhihuPinDaily)
		group.GET("/pin/hotlist", zhihuController.GetZhihuPinHotList)
		group.GET("/pin/people/:id", zhihuController.GetZhihuPinPeople)
		group.GET("/topic/:id", zhihuController.GetTopic)
	})
}
