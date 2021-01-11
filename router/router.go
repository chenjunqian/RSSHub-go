package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/router/routers"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Group("/zhihu", routers.ZhihubRouter)
		group.Group("/bilibili", routers.BilibiliRouter)
		group.Group("/bing", routers.BingRouter)
		group.Group("/weibo", routers.WeiboRouter)
		group.Group("/199IT", routers.IT199Router)
	})
}
