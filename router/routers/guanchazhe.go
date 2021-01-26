package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/guanchazhe"
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
