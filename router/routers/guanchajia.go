package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/guanchajia"
)

func GuanChaJiaRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		guanchajiaCtl := new(guanchajia.Controller)
		group.Group("/index", func(group *ghttp.RouterGroup) {
			group.GET("/shangyechanye", guanchajiaCtl.GetIndex)
			group.GET("/caijing", guanchajiaCtl.GetIndex)
			group.GET("/dichan", guanchajiaCtl.GetIndex)
			group.GET("/qiche", guanchajiaCtl.GetIndex)
			group.GET("/guanchajia", guanchajiaCtl.GetIndex)
			group.GET("/zhuanlan", guanchajiaCtl.GetIndex)
			group.GET("/lishi", guanchajiaCtl.GetIndex)
			group.GET("/shuping", guanchajiaCtl.GetIndex)
			group.GET("/zongshen", guanchajiaCtl.GetIndex)
			group.GET("/wenhua", guanchajiaCtl.GetIndex)
			group.GET("/lingdu", guanchajiaCtl.GetIndex)
		})
	})
}
