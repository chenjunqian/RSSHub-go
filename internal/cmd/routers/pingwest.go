package routers

import (
	"rsshub/internal/controller/rssapi/pingwest"

	"github.com/gogf/gf/v2/net/ghttp"
)

func PingwestRouter(group *ghttp.RouterGroup) {
	pingwestCtl := new(pingwest.Controller)
	group.GET("/state", pingwestCtl.GetState)
}
