package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/dongqiudi"
)

func DongQiuDiRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		dqdCtl := new(dongqiudi.Controller)
		group.GET("/daily", dqdCtl.GetDaily)
		group.GET("topnews/toutiao", dqdCtl.GetTopNews)
		group.GET("topnews/shendu", dqdCtl.GetTopNews)
		group.GET("topnews/xianqing", dqdCtl.GetTopNews)
		group.GET("topnews/dzhan", dqdCtl.GetTopNews)
		group.GET("topnews/zhongchao", dqdCtl.GetTopNews)
		group.GET("topnews/guoji", dqdCtl.GetTopNews)
		group.GET("topnews/yingchao", dqdCtl.GetTopNews)
		group.GET("topnews/xijia", dqdCtl.GetTopNews)
		group.GET("topnews/yijia", dqdCtl.GetTopNews)
		group.GET("topnews/dejia", dqdCtl.GetTopNews)
		group.GET("special/xinwendabaozha", dqdCtl.GetSpecial)
		group.GET("special/shijiaqiu", dqdCtl.GetSpecial)
		group.GET("special/mvp", dqdCtl.GetSpecial)
	})
}
