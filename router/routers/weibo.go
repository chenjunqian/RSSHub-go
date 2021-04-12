package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/weibo"
)

func WeiboRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		weiboController := new(weibo.Controller)
		group.GET("/search/hot/", weiboController.GetSearchHot)
	})
}
