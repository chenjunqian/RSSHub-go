package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/idaily"
)

func IDailyRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		iDailyCtl := new(idaily.Controller)
		group.GET("/index", iDailyCtl.GetIndex)
	})
}
