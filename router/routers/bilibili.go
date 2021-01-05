package routers

import (
	"github.com/gogf/gf/net/ghttp"

	"rsshub/app/api/bilibili"
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
	})
}
