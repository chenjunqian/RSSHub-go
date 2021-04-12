package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/sciencenet"
)

func ScienceNetRouter(group *ghttp.RouterGroup) {
	group.Group("/index", func(group *ghttp.RouterGroup) {
		sciencenetCtl := new(sciencenet.Controller)
		group.GET("/recommend", sciencenetCtl.GetIndex)
		group.GET("/hot", sciencenetCtl.GetIndex)
		group.GET("/new", sciencenetCtl.GetIndex)
	})
}
