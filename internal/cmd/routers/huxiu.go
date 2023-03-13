package routers

import (
	"rsshub/internal/controller/rssapi/huxiu"

	"github.com/gogf/gf/v2/net/ghttp"
)

func HuXiuRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		huxiuController := new(huxiu.Controller)
		group.GET("/article", huxiuController.GetArticle)
		group.GET("/event", huxiuController.GetEvent)
		group.GET("/timeline", huxiuController.GetTimeline)
		group.GET("/collection", huxiuController.GetCollection)
		group.Group("/channel", func(group *ghttp.RouterGroup) {
			group.GET("/auto", huxiuController.GetChannels)
			group.GET("/young", huxiuController.GetChannels)
			group.GET("/consumer", huxiuController.GetChannels)
			group.GET("/tech", huxiuController.GetChannels)
			group.GET("/finance", huxiuController.GetChannels)
			group.GET("/entertainment", huxiuController.GetChannels)
			group.GET("/medical", huxiuController.GetChannels)
			group.GET("/culture", huxiuController.GetChannels)
			group.GET("/oversea", huxiuController.GetChannels)
			group.GET("/realestate", huxiuController.GetChannels)
			group.GET("/enterprise", huxiuController.GetChannels)
			group.GET("/startup", huxiuController.GetChannels)
			group.GET("/social", huxiuController.GetChannels)
			group.GET("/global", huxiuController.GetChannels)
			group.GET("/life", huxiuController.GetChannels)
		})
	})
}
