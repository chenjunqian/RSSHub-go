package routers

import (
	"rsshub/app/api/rssapi/niaogenote"

	"github.com/gogf/gf/v2/net/ghttp"
)

func NiaogeNoteRouter(group *ghttp.RouterGroup) {
	group.Group("/cat", func(group *ghttp.RouterGroup) {
		niaogeCtl := new(niaogenote.Controller)
		group.GET("/user_op", niaogeCtl.GetCat)
		group.GET("/activity_op", niaogeCtl.GetCat)
		group.GET("/new_media", niaogeCtl.GetCat)
		group.GET("/data_op", niaogeCtl.GetCat)
		group.GET("/video_live", niaogeCtl.GetCat)
		group.GET("/e_fast_selling", niaogeCtl.GetCat)
		group.GET("/ASO", niaogeCtl.GetCat)
		group.GET("/SEM", niaogeCtl.GetCat)
		group.GET("/info_stream", niaogeCtl.GetCat)
		group.GET("/marking_promotion", niaogeCtl.GetCat)
		group.GET("/brand_strategy", niaogeCtl.GetCat)
		group.GET("/ad", niaogeCtl.GetCat)
		group.GET("/create_activity", niaogeCtl.GetCat)
		group.GET("/career_growth", niaogeCtl.GetCat)
		group.GET("/product_design", niaogeCtl.GetCat)
		group.GET("/eff_tool", niaogeCtl.GetCat)
		group.GET("/management", niaogeCtl.GetCat)
	})
}
