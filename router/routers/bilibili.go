package routers

import (
	"github.com/gogf/gf/net/ghttp"

	"rsshub/app/api/rssapi/bilibili"
)

func BilibiliRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		bilibiliController := new(bilibili.Controller)
		group.GET("/appversion/:id", bilibiliController.GetAppVersion)
		group.GET("/user/article/:id", bilibiliController.GetUserArticle)
		group.GET("/user/audio/:id", bilibiliController.GetUserAudio)
		group.GET("/bangumi/:id", bilibiliController.GetBangumi)
		group.GET("/blackboard", bilibiliController.GetBlackboard)
		group.GET("/linkNews/:product", bilibiliController.GetLinkNews)
		group.GET("/live/area/:areaId/:order", bilibiliController.GetLinvArea)
		group.GET("/live/room/:roomId", bilibiliController.GetLinkRoom)
		group.GET("/manga/update/:id", bilibiliController.GetMangaUpdate)
		group.GET("/readlist/:id", bilibiliController.GetReadList)
		group.GET("/topic/:topicName", bilibiliController.GetTopic)
		group.GET("/user/fav/:id", bilibiliController.GetUserFav)
		group.GET("/weekly", bilibiliController.GetWeeklyRecommend)
	})
}
