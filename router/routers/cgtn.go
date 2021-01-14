package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/cgtn"
)

func CGTNRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		cgtnController := new(cgtn.Controller)
		group.GET("/most/all", cgtnController.GetMostRead)
		group.GET("/most/day", cgtnController.GetMostRead)
		group.GET("/most/week", cgtnController.GetMostRead)
		group.GET("/most/month", cgtnController.GetMostRead)
		group.GET("/most/year", cgtnController.GetMostRead)
		group.GET("/top", cgtnController.GetTop)
	})
}
