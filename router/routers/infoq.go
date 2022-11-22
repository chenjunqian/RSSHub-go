package routers

import (
	"rsshub/app/api/rssapi/infoq"

	"github.com/gogf/gf/v2/net/ghttp"
)

func InfoQRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		infoqCtl := new(infoq.Controller)
		group.GET("/recommend", infoqCtl.GetRecommend)
	})
}
