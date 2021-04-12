package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/pingwest"
)

func PingwestRouter(group *ghttp.RouterGroup) {
	pingwestCtl := new(pingwest.Controller)
	group.GET("/state", pingwestCtl.GetState)
}
