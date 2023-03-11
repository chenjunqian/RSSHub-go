package routers

import (
	"rsshub/internal/controller/webapi"

	"github.com/gogf/gf/v2/net/ghttp"
)

func APIRouter(group *ghttp.RouterGroup) {
	group.Group("/v1", func(group *ghttp.RouterGroup) {
		//group.Middleware(middleware.AuthToken)
		webController := new(webapi.Controller)
		group.Group("", func(group *ghttp.RouterGroup) {
			group.GET("/routers", webController.GetAllRssResource)
			group.GET("/feed/info/list", webController.GetAllFeedChannelInfoList)
		})
	})
}

func WebRouter(group *ghttp.RouterGroup) {
	webController := new(webapi.Controller)
	group.GET("/", webController.IndexTpl)
	group.GET("/s/:keyword/:start", webController.SearchFeedItems)
	group.GET("/f/c/:id", webController.FeedChannelDetail)
	group.GET("/f/i/:id", webController.FeedItemDetail)
}
