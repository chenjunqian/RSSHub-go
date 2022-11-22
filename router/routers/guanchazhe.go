package routers

import (
	"rsshub/app/api/rssapi/guanchazhe"

	"github.com/gogf/gf/v2/net/ghttp"
)

func GuanChaZheRouter(group *ghttp.RouterGroup) {
	group.Group("", func(group *ghttp.RouterGroup) {
		guanchazheCtl := new(guanchazhe.Controller)
		group.GET("/headline", guanchazheCtl.GetHeadLine)
		group.GET("/internation", guanchazheCtl.GetIndex)
		group.GET("/military", guanchazheCtl.GetIndex)
		group.GET("/economy", guanchazheCtl.GetIndex)
		group.GET("/tech", guanchazheCtl.GetIndex)
		group.GET("/auto", guanchazheCtl.GetIndex)
	})
}
