package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/guokr"
)

func GuoKrRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		guokrController := new(guokr.Controller)
		group.Group("/science", func(group *ghttp.RouterGroup) {
			group.Group("/category", func(group *ghttp.RouterGroup) {
				group.GET("/science/", guokrController.GetScienceCategory)
				group.GET("/funny/", guokrController.GetScienceCategory)
				group.GET("/life/", guokrController.GetScienceCategory)
				group.GET("/health/", guokrController.GetScienceCategory)
				group.GET("/humanities/", guokrController.GetScienceCategory)
				group.GET("/nature/", guokrController.GetScienceCategory)
				group.GET("/digital/", guokrController.GetScienceCategory)
				group.GET("/food/", guokrController.GetScienceCategory)
			})
			group.Group("/subject", func(group *ghttp.RouterGroup) {
				group.GET("/engineering/", guokrController.GetScienceSubject)
				group.GET("/education/", guokrController.GetScienceSubject)
				group.GET("/physics/", guokrController.GetScienceSubject)
				group.GET("/sex/", guokrController.GetScienceSubject)
				group.GET("/agronomy/", guokrController.GetScienceSubject)
				group.GET("/psychology/", guokrController.GetScienceSubject)
				group.GET("/medicine/", guokrController.GetScienceSubject)
				group.GET("/forensic/", guokrController.GetScienceSubject)
				group.GET("/society/", guokrController.GetScienceSubject)
				group.GET("/atmosphere/", guokrController.GetScienceSubject)
				group.GET("/others/", guokrController.GetScienceSubject)
				group.GET("/chemistry/", guokrController.GetScienceSubject)
				group.GET("/earth/", guokrController.GetScienceSubject)
				group.GET("/communication/", guokrController.GetScienceSubject)
				group.GET("/environment/", guokrController.GetScienceSubject)
				group.GET("/diy/", guokrController.GetScienceSubject)
				group.GET("/astronomy/", guokrController.GetScienceSubject)
				group.GET("/math/", guokrController.GetScienceSubject)
				group.GET("/biology/", guokrController.GetScienceSubject)
				group.GET("/aerospace/", guokrController.GetScienceSubject)
				group.GET("/internet/", guokrController.GetScienceSubject)
			})
		})
		group.GET("/calendar/", guokrController.GetIndex)
		group.GET("/foodlab/", guokrController.GetIndex)
		group.GET("/pretty/", guokrController.GetIndex)
	})
}
