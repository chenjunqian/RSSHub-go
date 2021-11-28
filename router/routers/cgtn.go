package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/cgtn"
)

func CGTNRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/most/all", cgtn.Controller.GetMostRead)
		group.GET("/most/day", cgtn.Controller.GetMostRead)
		group.GET("/most/week", cgtn.Controller.GetMostRead)
		group.GET("/most/month", cgtn.Controller.GetMostRead)
		group.GET("/most/year", cgtn.Controller.GetMostRead)
		group.GET("/top", cgtn.Controller.GetTop)
	})
}
