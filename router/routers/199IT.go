package routers

import (
	"github.com/gogf/gf/net/ghttp"
	_199IT "rsshub/app/api/rssapi/199IT"
)

func IT199Router(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/index", _199IT.IT1999Controller.Get199ITIndex)
		group.GET("/category/report", _199IT.IT1999Controller.Get199ITCategoryReport)
	})
}
