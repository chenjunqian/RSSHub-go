package routers

import (
	"rsshub/app/api/rssapi/baidu"

	"github.com/gogf/gf/v2/net/ghttp"
)

func BaiduRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/zhidao/daily", baidu.BaiDuController.GetZhiDaoDaily)
	})
}
