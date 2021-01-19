package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/cyzone"
)

func CYZoneRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		cyZoneCtl := new(cyzone.Controller)
		group.GET("/news", cyZoneCtl.GetNews)
		group.GET("/kuailiyu", cyZoneCtl.GetNews)
		group.GET("/chuangtou", cyZoneCtl.GetNews)
		group.GET("/kechuang", cyZoneCtl.GetNews)
		group.GET("/qiche", cyZoneCtl.GetNews)
		group.GET("/haiwai", cyZoneCtl.GetNews)
		group.GET("/xiaofei", cyZoneCtl.GetNews)
		group.GET("/keji", cyZoneCtl.GetNews)
		group.GET("/yiliao", cyZoneCtl.GetNews)
		group.GET("/wenyu", cyZoneCtl.GetNews)
		group.GET("/chengshi", cyZoneCtl.GetNews)
		group.GET("/zhengce", cyZoneCtl.GetNews)
		group.GET("/texie", cyZoneCtl.GetNews)
		group.GET("/ganhuo", cyZoneCtl.GetNews)
		group.GET("/kejigu", cyZoneCtl.GetNews)
	})
}
