package routers

import (
	"rsshub/internal/controller/rssapi/weibo"

	"github.com/gogf/gf/v2/net/ghttp"
)

func WeiboRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		weiboController := new(weibo.Controller)
		group.GET("/search/hot/", weiboController.GetSearchHot)
	})
}
