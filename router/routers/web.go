package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/webapi"
)

func WebRouter(group *ghttp.RouterGroup) {
	group.Group("/v1", func(group *ghttp.RouterGroup) {
		//group.Middleware(middleware.AuthToken)
		webController := new(webapi.Controller)
		group.Group("", func(group *ghttp.RouterGroup) {
			group.GET("/routers", webController.GetAllRssResource)
			group.GET("/feed_tag", webController.GetFeedTag)
			group.GET("/feed_channel_by_tag", webController.GetFeedChannelByTag)
			group.GET("/feed_item_by_channel_id", webController.GetFeedItemByChannelId)
		})
	})
}
