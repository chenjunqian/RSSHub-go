package baijing

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) GetEBusiness(req *ghttp.Request) {
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
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := commonHtmlParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", "BAIJING_EBUSINESS", rssStr)
	g.Redis().DoVar("EXPIRE", "BAIJING_EBUSINESS", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
