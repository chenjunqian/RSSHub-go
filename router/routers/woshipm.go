package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/woshipm"
)

func WoshipmRouter(group *ghttp.RouterGroup) {
	woshipmCtl := new(woshipm.Controller)
	group.GET("/latest", woshipmCtl.GetIndex)
	group.GET("/popular", woshipmCtl.GetPopular)
}
