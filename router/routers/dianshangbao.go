package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/dianshangbao"
)

func DSBRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/lingshou", dianshangbao.Controller.GetIndex)
		group.GET("/wuliu", dianshangbao.Controller.GetIndex)
		group.GET("/O2O", dianshangbao.Controller.GetIndex)
		group.GET("/zhifu", dianshangbao.Controller.GetIndex)
		group.GET("/B2B", dianshangbao.Controller.GetIndex)
		group.GET("/renwu", dianshangbao.Controller.GetIndex)
		group.GET("/kuajing", dianshangbao.Controller.GetIndex)
		group.GET("/guancha", dianshangbao.Controller.GetIndex)
	})
}
