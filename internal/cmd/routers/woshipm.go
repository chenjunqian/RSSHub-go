package routers

import (
	"rsshub/internal/controller/rssapi/woshipm"

	"github.com/gogf/gf/v2/net/ghttp"
)

func WoshipmRouter(group *ghttp.RouterGroup) {
	woshipmCtl := new(woshipm.Controller)
	group.GET("/latest", woshipmCtl.GetIndex)
	group.GET("/popular", woshipmCtl.GetPopular)
}
