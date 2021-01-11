package routers

import (
	"github.com/gogf/gf/net/ghttp"
	_199IT "rsshub/app/api/199IT"
)

func IT199Router(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		it199Controller := new(_199IT.Controller)
		group.GET("/index", it199Controller.Get199ITIndex)
		group.GET("/category/report", it199Controller.Get199ITCategoryReport)
	})
}
