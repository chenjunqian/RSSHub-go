package routers

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/api/rssapi/ifeng"
)

func IFengRouter(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		ifengController := new(ifeng.Controller)
		group.Group("/news", func(group *ghttp.RouterGroup) {
			group.GET("/living/", ifengController.GetNews)
			group.GET("/bigfish/", ifengController.GetNews)
			group.GET("/globalyouth/", ifengController.GetNews)
			group.GET("/warmstory/", ifengController.GetNews)
			group.GET("/fhphotos/", ifengController.GetNews)
			group.GET("/xuanzhan2020/", ifengController.GetNews)
		})
		group.Group("/finance", func(group *ghttp.RouterGroup) {
			group.GET("/hk_hotpoint/", ifengController.GetFinance)
			group.GET("/hk_baishitong/", ifengController.GetFinance)
			group.GET("/hk_xiaozhishi/", ifengController.GetFinance)
			group.GET("/globalyouth/", ifengController.GetFinance)
			group.GET("/hk_zhuanlan/", ifengController.GetFinance)
			group.GET("/hk_jigou_dongtai/", ifengController.GetFinance)
			group.GET("/hk_tianxia/", ifengController.GetFinance)
			group.GET("/ipo_kechuang/", ifengController.GetFinance)
		})
		group.Group("/money", func(group *ghttp.RouterGroup) {
			group.GET("/hot/", ifengController.GetMoney)
			group.GET("/bank/", ifengController.GetMoney)
			group.GET("/insure/", ifengController.GetMoney)
			group.GET("/fund/", ifengController.GetMoney)
		})
		group.Group("/ent", func(group *ghttp.RouterGroup) {
			group.GET("/star/", ifengController.GetEntertainment)
			group.GET("/movie/", ifengController.GetEntertainment)
			group.GET("/tv/", ifengController.GetEntertainment)
			group.GET("/music/", ifengController.GetEntertainment)
		})
		group.Group("/culture", func(group *ghttp.RouterGroup) {
			group.GET("/read/", ifengController.GetCulture)
			group.GET("/artist/", ifengController.GetCulture)
			group.GET("/insight/", ifengController.GetCulture)
			group.GET("/news/", ifengController.GetCulture)
		})
		group.Group("/fashion", func(group *ghttp.RouterGroup) {
			group.GET("/trends/", ifengController.GetFashion)
			group.GET("/beauty/", ifengController.GetFashion)
			group.GET("/lifestyle/", ifengController.GetFashion)
			group.GET("/emotion/", ifengController.GetFashion)
		})
		group.Group("/auto", func(group *ghttp.RouterGroup) {
			group.GET("/xinche/", ifengController.GetAuto)
			group.GET("/shijia/", ifengController.GetAuto)
			group.GET("/daogou/", ifengController.GetAuto)
			group.GET("/hangye/", ifengController.GetAuto)
		})
		group.Group("/tech", func(group *ghttp.RouterGroup) {
			group.GET("/index/", ifengController.GetTech)
			group.GET("/digi/", ifengController.GetTech)
			group.GET("/mobile/", ifengController.GetTech)
			group.GET("/hangye/", ifengController.GetTech)
		})
	})
}
