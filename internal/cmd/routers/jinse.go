package routers

import (
	"rsshub/internal/controller/rssapi/jinse"

	"github.com/gogf/gf/v2/net/ghttp"
)

func JinseRouter(group *ghttp.RouterGroup) {
	jinseCtl := new(jinse.Controller)
	group.Group("/catalogue", func(group *ghttp.RouterGroup) {
		group.GET("/zhengce", jinseCtl.GetCatalogue)
		group.GET("/fenxishishuo", jinseCtl.GetCatalogue)
		group.GET("/defi", jinseCtl.GetCatalogue)
		group.GET("/kuang", jinseCtl.GetCatalogue)
		group.GET("/industry", jinseCtl.GetCatalogue)
		group.GET("/IPFS", jinseCtl.GetCatalogue)
		group.GET("/tech", jinseCtl.GetCatalogue)
		group.GET("/baike", jinseCtl.GetCatalogue)
		group.GET("/capitalmarket", jinseCtl.GetCatalogue)
	})
	group.GET("/lives", jinseCtl.GetLives)
	group.GET("/timeline", jinseCtl.GetTimeline)
}
