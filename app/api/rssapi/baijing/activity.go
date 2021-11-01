package baijing

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) GetActivity(req *ghttp.Request) {
	if value, err := g.Redis().DoVar("GET", "BAIJING_ACTIVITY"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.baijingapp.com/index/ajax/getlist/type-5"
	rssData := dao.RSSFeed{
		Title:       "白鲸出海-活动",
		Link:        "https://www.baijingapp.com",
		Description: "白鲸出海活动",
		Tag:         []string{"互联网", "开发者", "科技", "社区"},
		ImageUrl:    "https://www.baijingapp.com/static/css/default/img/favicon.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := commonHtmlParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "BAIJING_ACTIVITY", rssStr)
	g.Redis().DoVar("EXPIRE", "BAIJING_ACTIVITY", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
