package routers

import (
	"github.com/gogf/gf/net/ghttp"
	_36kr "rsshub/app/api/36kr"
)

func Kr36Router(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		kr36Controller := new(_36kr.Controller)
		group.GET("/news/latest", kr36Controller.Get36krNews)
		group.GET("/news/recommend", kr36Controller.Get36krNews)
		group.GET("/news/contact", kr36Controller.Get36krNews)
		group.GET("/news/ccs", kr36Controller.Get36krNews)
		group.GET("/news/travel", kr36Controller.Get36krNews)
		group.GET("/news/technology", kr36Controller.Get36krNews)
		group.GET("/news/enterpriseservice", kr36Controller.Get36krNews)
		group.GET("/news/banking", kr36Controller.Get36krNews)
		group.GET("/news/life", kr36Controller.Get36krNews)
		group.GET("/news/innovate", kr36Controller.Get36krNews)
		group.GET("/news/estate", kr36Controller.Get36krNews)
		group.GET("/news/workplace", kr36Controller.Get36krNews)
		group.GET("/news/other", kr36Controller.Get36krNews)
		group.GET("/news/flashes", kr36Controller.Get36krNewsFlashes)
		group.GET("/user/:id", kr36Controller.Get36krUserNews)
	})
}
