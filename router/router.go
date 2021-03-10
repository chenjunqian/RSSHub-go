package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/router/routers"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Group("/zhihu", routers.ZhihubRouter)
		group.Group("/bilibili", routers.BilibiliRouter)
		group.Group("/bing", routers.BingRouter)
		group.Group("/weibo", routers.WeiboRouter)
		group.Group("/199IT", routers.IT199Router)
		group.Group("/36kr", routers.Kr36Router)
		group.Group("/cgtn", routers.CGTNRouter)
		group.Group("/cnbeta", routers.CNBetaRouter)
		group.Group("/dayone", routers.DayOneRouter)
		group.Group("/engadget", routers.EngadgetRouter)
		group.Group("/idaily", routers.IDailyRouter)
		group.Group("/infoq", routers.InfoQRouter)
		group.Group("/mitchina", routers.MitChinaRouter)
		group.Group("/ifan", routers.IFanRouter)
		group.Group("/baidu", routers.BaiduRouter)
		group.Group("/baijing", routers.BaijingRouter)
		group.Group("/bishijie", routers.BiShiJieRouter)
		group.Group("/chaping", routers.ChaPingRouter)
		group.Group("/chouti", routers.ChouTiRouter)
		group.Group("/cyzone", routers.CYZoneRouter)
		group.Group("/dianshangbao", routers.DSBRouter)
		group.Group("/dongqiudi", routers.DongQiuDiRouter)
		group.Group("/dx2025", routers.DX2025Router)
		group.Group("/duozhi", routers.DuoZhiRouter)
		group.Group("/ifeng", routers.IFengRouter)
		group.Group("/fulinian", routers.FuLiNianRouter)
		group.Group("/guanchajia", routers.GuanChaJiaRouter)
		group.Group("/guanchazhe", routers.GuanChaZheRouter)
		group.Group("/guokr", routers.GuoKrRouter)
		group.Group("/houxu", routers.HouXuRouter)
		group.Group("/huxiu", routers.HuXiuRouter)
		group.Group("/gcore", routers.GCoreRouter)
		group.Group("/jinse", routers.JinseRouter)
		group.Group("/whalegogo", routers.WhaleGoGoRouter)
		group.Group("/juesheng", routers.JueShengRouter)
		group.Group("/sciencenet", routers.ScienceNetRouter)
		group.Group("/meihua", routers.MeiHuaRouter)
		group.Group("/medsci", routers.MedsciRouter)
		group.Group("/niaogebiji", routers.NiaogeNoteRouter)
		group.Group("/pintu", routers.PintuRouter)
	})
}
