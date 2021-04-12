package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/dianshangbao"
)

func DSBRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		dsbCtl := new(dianshangbao.Controller)
		group.GET("/lingshou", dsbCtl.GetIndex)
		group.GET("/wuliu", dsbCtl.GetIndex)
		group.GET("/O2O", dsbCtl.GetIndex)
		group.GET("/zhifu", dsbCtl.GetIndex)
		group.GET("/B2B", dsbCtl.GetIndex)
		group.GET("/renwu", dsbCtl.GetIndex)
		group.GET("/kuajing", dsbCtl.GetIndex)
		group.GET("/guancha", dsbCtl.GetIndex)
	})
}
