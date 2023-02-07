package routers

import (
	"rsshub/internal/controller/rssapi/duozhi"

	"github.com/gogf/gf/v2/net/ghttp"
)

func DuoZhiRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		duozhiCtl := new(duozhi.Controller)
		group.GET("/industry/insight", duozhiCtl.GetIndustryNews)
		group.GET("/industry/preschool", duozhiCtl.GetIndustryNews)
		group.GET("/industry/K12", duozhiCtl.GetIndustryNews)
		group.GET("/industry/qualityedu", duozhiCtl.GetIndustryNews)
		group.GET("/industry/adultedu", duozhiCtl.GetIndustryNews)
		group.GET("/industry/EduInformatization", duozhiCtl.GetIndustryNews)
		group.GET("/industry/earnings", duozhiCtl.GetIndustryNews)
		group.GET("/industry/privateschools", duozhiCtl.GetIndustryNews)
		group.GET("/industry/overseas", duozhiCtl.GetIndustryNews)
	})
}
