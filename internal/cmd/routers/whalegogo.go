package routers

import (
	"rsshub/internal/controller/rssapi/whalegogo"

	"github.com/gogf/gf/v2/net/ghttp"
)

func WhaleGoGoRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		whalegogoCtl := new(whalegogo.Controller)
		group.GET("/index", whalegogoCtl.GetIndex)

		group.Group("/portal", func(group *ghttp.RouterGroup) {
			group.GET("/news", whalegogoCtl.GetIPortal)
			group.GET("/article", whalegogoCtl.GetIPortal)
			group.GET("/activities", whalegogoCtl.GetIPortal)
			group.GET("/appraisal", whalegogoCtl.GetIPortal)
		})
	})
}
