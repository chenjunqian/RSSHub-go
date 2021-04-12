package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/whalegogo"
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
