package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/infoq"
)

func InfoQRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		infoqCtl := new(infoq.Controller)
		group.GET("/recommend", infoqCtl.GetRecommend)
	})
}
