package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/baijing"
)

func BaijingRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		baijingCtl := new(baijing.Controller)
		group.GET("/daily", baijingCtl.GetNews)
		group.GET("/weekly", baijingCtl.GetWeekly)
		group.GET("/ganhuo", baijingCtl.GetGanHuo)
		group.GET("/zhuanlan", baijingCtl.GetZhuanlan)
		group.GET("/shouyou", baijingCtl.GetShouyou)
		group.GET("/touzi", baijingCtl.GetTouzi)
		group.GET("/datareport", baijingCtl.GetDataReport)
		group.GET("/mobile", baijingCtl.GetMobilePhone)
		group.GET("/ebusiness", baijingCtl.GetEBusiness)
		group.GET("/activity", baijingCtl.GetActivity)
	})
}
