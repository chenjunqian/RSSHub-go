package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/baidu"
)

func BaiduRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/zhidao/daily", baidu.BaiDuController.GetZhiDaoDaily)
	})
}
