package routers

import (
	"rsshub/app/api/rssapi/juejin"

	"github.com/gogf/gf/net/ghttp"
)

func JuejinRouter(group *ghttp.RouterGroup) {
	group.Group("/recommand", func(group *ghttp.RouterGroup) {
		group.GET("/hot", juejin.Controller.GetRecommand)
	})
}