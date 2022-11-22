package routers

import (
	"rsshub/app/api/rssapi/baijing"

	"github.com/gogf/gf/v2/net/ghttp"
)

func BaijingRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/daily", baijing.BJController.GetNews)
		group.GET("/weekly", baijing.BJController.GetWeekly)
		group.GET("/ganhuo", baijing.BJController.GetGanHuo)
		group.GET("/zhuanlan", baijing.BJController.GetZhuanlan)
		group.GET("/shouyou", baijing.BJController.GetShouyou)
		group.GET("/touzi", baijing.BJController.GetTouzi)
		group.GET("/datareport", baijing.BJController.GetDataReport)
		group.GET("/mobile", baijing.BJController.GetMobilePhone)
		group.GET("/ebusiness", baijing.BJController.GetEBusiness)
		group.GET("/activity", baijing.BJController.GetActivity)
	})
}
