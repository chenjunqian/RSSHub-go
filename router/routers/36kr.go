package routers

import (
	"github.com/gogf/gf/net/ghttp"
	_36kr "rsshub/app/api/rssapi/36kr"
)

func Kr36Router(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/news/latest", _36kr.KR36Controller.Get36krNews)
		group.GET("/news/recommend", _36kr.KR36Controller.Get36krNews)
		group.GET("/news/contact", _36kr.KR36Controller.Get36krNews)
		group.GET("/news/ccs", _36kr.KR36Controller.Get36krNews)
		group.GET("/news/travel", _36kr.KR36Controller.Get36krNews)
		group.GET("/news/technology", _36kr.KR36Controller.Get36krNews)
		group.GET("/news/enterpriseservice", _36kr.KR36Controller.Get36krNews)
		group.GET("/news/banking", _36kr.KR36Controller.Get36krNews)
		group.GET("/news/life", _36kr.KR36Controller.Get36krNews)
		group.GET("/news/innovate", _36kr.KR36Controller.Get36krNews)
		group.GET("/news/estate", _36kr.KR36Controller.Get36krNews)
		group.GET("/news/workplace", _36kr.KR36Controller.Get36krNews)
		group.GET("/news/other", _36kr.KR36Controller.Get36krNews)
		group.GET("/news/flashes", _36kr.KR36Controller.Get36krNewsFlashes)
		group.GET("/user/:id", _36kr.KR36Controller.Get36krUserNews)
	})
}
