package routers

import (
	"rsshub/internal/controller/rssapi/cyzone"

	"github.com/gogf/gf/v2/net/ghttp"
)

func CYZoneRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/news", cyzone.Controller.GetNews)
		group.GET("/kuailiyu", cyzone.Controller.GetNews)
		group.GET("/chuangtou", cyzone.Controller.GetNews)
		group.GET("/kechuang", cyzone.Controller.GetNews)
		group.GET("/qiche", cyzone.Controller.GetNews)
		group.GET("/haiwai", cyzone.Controller.GetNews)
		group.GET("/xiaofei", cyzone.Controller.GetNews)
		group.GET("/keji", cyzone.Controller.GetNews)
		group.GET("/yiliao", cyzone.Controller.GetNews)
		group.GET("/wenyu", cyzone.Controller.GetNews)
		group.GET("/chengshi", cyzone.Controller.GetNews)
		group.GET("/zhengce", cyzone.Controller.GetNews)
		group.GET("/texie", cyzone.Controller.GetNews)
		group.GET("/ganhuo", cyzone.Controller.GetNews)
		group.GET("/kejigu", cyzone.Controller.GetNews)
	})
}
