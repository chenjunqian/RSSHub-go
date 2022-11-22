package routers

import (
	"rsshub/app/api/rssapi/dx2025"

	"github.com/gogf/gf/v2/net/ghttp"
)

func DX2025Router(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		dx2025Ctl := new(dx2025.Controller)
		group.GET("report/new-tch-info", dx2025Ctl.GetCategoryNews)
		group.GET("report/robot-idt-info", dx2025Ctl.GetCategoryNews)
		group.GET("report/aerospace-info", dx2025Ctl.GetCategoryNews)
		group.GET("report/marine-info", dx2025Ctl.GetCategoryNews)
		group.GET("report/transportation-info", dx2025Ctl.GetCategoryNews)
		group.GET("report/energy-vehicles-info", dx2025Ctl.GetCategoryNews)
		group.GET("report/electric-equipment-info", dx2025Ctl.GetCategoryNews)
		group.GET("report/agricultural-equipment-info", dx2025Ctl.GetCategoryNews)
		group.GET("report/material-info", dx2025Ctl.GetCategoryNews)
		group.GET("report/biomedicine-medical-info", dx2025Ctl.GetCategoryNews)
		group.GET("report/modern-service-info", dx2025Ctl.GetCategoryNews)
		group.GET("report/manufacturing-info", dx2025Ctl.GetCategoryNews)

		group.GET("observation/new-tch-info", dx2025Ctl.GetCategoryNews)
		group.GET("observation/robot-idt-info", dx2025Ctl.GetCategoryNews)
		group.GET("observation/aerospace-info", dx2025Ctl.GetCategoryNews)
		group.GET("observation/marine-info", dx2025Ctl.GetCategoryNews)
		group.GET("observation/transportation-info", dx2025Ctl.GetCategoryNews)
		group.GET("observation/energy-vehicles-info", dx2025Ctl.GetCategoryNews)
		group.GET("observation/electric-equipment-info", dx2025Ctl.GetCategoryNews)
		group.GET("observation/agricultural-equipment-info", dx2025Ctl.GetCategoryNews)
		group.GET("observation/material-info", dx2025Ctl.GetCategoryNews)
		group.GET("observation/biomedicine-medical-info", dx2025Ctl.GetCategoryNews)
		group.GET("observation/modern-service-info", dx2025Ctl.GetCategoryNews)
		group.GET("observation/manufacturing-info", dx2025Ctl.GetCategoryNews)
	})
}
