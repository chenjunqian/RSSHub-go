package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/medsci"
)

func MedsciRouter(group *ghttp.RouterGroup) {
	group.Group("/index", func(group *ghttp.RouterGroup) {
		medsciCtl := new(medsci.Controller)
		group.GET("/recommend", medsciCtl.GetIndex)
	})
}
