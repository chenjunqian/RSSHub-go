package whalegogo

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {
	cacheKey := "WHALE_GOGO_INDEX"
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://api.whalegogo.com/v1/app/index"
	rssData := dao.RSSFeed{
		Title:       "鲸跃汽车 - 最新",
		Tag:         []string{"汽车"},
		Link:        "https://m.whalegogo.com/index",
		Description: "我们是一帮在传统汽车门户、汽车杂志战斗过数年的老司机，关于车与理想生活的种种，我们有很多思考，现在用一个全新的网站和 APP，将车型快讯、试驾体验、行业深度、生活方式等原创内容，用简洁的设计与排版呈现在你的面前。",
		ImageUrl:    "https://api.whalegogo.com/favicon.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := indexParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
