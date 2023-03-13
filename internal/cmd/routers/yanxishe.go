package routers

import (
	"rsshub/internal/controller/rssapi/yanxishe"

	"github.com/gogf/gf/v2/net/ghttp"
)

func YanXiSheRouter(group *ghttp.RouterGroup) {
	group.GET("/index", yanxishe.Controller.GetIndex)
}
