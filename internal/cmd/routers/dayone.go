package routers

import (
	"rsshub/internal/controller/rssapi/dayone"

	"github.com/gogf/gf/v2/net/ghttp"
)

func DayOneRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		dayoneCtl := new(dayone.Controller)
		group.GET("/blog", dayoneCtl.GetMostRead)
	})
}
