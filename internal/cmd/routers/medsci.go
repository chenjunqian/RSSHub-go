package routers

import (
	"rsshub/internal/controller/rssapi/medsci"

	"github.com/gogf/gf/v2/net/ghttp"
)

func MedsciRouter(group *ghttp.RouterGroup) {
	group.Group("/index", func(group *ghttp.RouterGroup) {
		medsciCtl := new(medsci.Controller)
		group.GET("/recommend", medsciCtl.GetIndex)
	})
}
