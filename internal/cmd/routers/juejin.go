package routers

import (
	"rsshub/internal/controller/rssapi/juejin"

	"github.com/gogf/gf/v2/net/ghttp"
)

func JuejinRouter(group *ghttp.RouterGroup) {
	group.Group("/recommand", func(group *ghttp.RouterGroup) {
		group.GET("/hot", juejin.Controller.GetRecommand)
	})
}
