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
		group.Group("/36kr", routers.Kr36Router)
		group.Group("/cgtn", routers.CGTNRouter)
		group.Group("/cnbeta", routers.CNBetaRouter)
		group.Group("/dayone", routers.DayOneRouter)
		group.Group("/engadget", routers.EngadgetRouter)
	})
}
