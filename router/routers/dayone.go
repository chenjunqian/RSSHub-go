package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/dayone"
)

func DayOneRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		dayoneCtl := new(dayone.Controller)
		group.GET("/blog", dayoneCtl.GetMostRead)
	})
}
