package routers

import (
	"rsshub/app/api/rssapi/sciencenet"

	"github.com/gogf/gf/v2/net/ghttp"
)

func ScienceNetRouter(group *ghttp.RouterGroup) {
	group.Group("/index", func(group *ghttp.RouterGroup) {
		sciencenetCtl := new(sciencenet.Controller)
		group.GET("/recommend", sciencenetCtl.GetIndex)
		group.GET("/hot", sciencenetCtl.GetIndex)
		group.GET("/new", sciencenetCtl.GetIndex)
	})
}
