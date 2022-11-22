package routers

import (
	"rsshub/app/api/rssapi/idaily"

	"github.com/gogf/gf/v2/net/ghttp"
)

func IDailyRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		iDailyCtl := new(idaily.Controller)
		group.GET("/index", iDailyCtl.GetIndex)
	})
}
