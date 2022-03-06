package routers

import (
	"rsshub/app/api/rssapi/yanxishe"

	"github.com/gogf/gf/net/ghttp"
)

func YanXiSheRouter(group *ghttp.RouterGroup) {
	group.GET("/index", yanxishe.Controller.GetIndex)
}