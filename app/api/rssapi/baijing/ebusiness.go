package baijing

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *controller) GetEBusiness(req *ghttp.Request) {
	if value, err := g.Redis().DoVar("GET", "BAIJING_EBUSINESS"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.baijingapp.com/index/ajax/getlist/type-5"
	rssData := dao.RSSFeed{
		Title:       "白鲸出海-电商",
		Link:        "https://www.baijingapp.com",
		Description: "白鲸出海电商",
		Tag:         []string{"电商", "数据"},
		ImageUrl:    "https://www.baijingapp.com/static/css/default/img/favicon.ico",
	}
	if resp := component.GetContent(apiUrl); resp != "" {
		rssItems := commonHtmlParser(resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "BAIJING_EBUSINESS", rssStr)
	g.Redis().DoVar("EXPIRE", "BAIJING_EBUSINESS", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
