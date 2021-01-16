package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/baidu"
)

func BaiduRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		baiduCtl := new(baidu.Controller)
		group.GET("/zhidao/daily", baiduCtl.GetZhiDaoDaily)
	})
}
