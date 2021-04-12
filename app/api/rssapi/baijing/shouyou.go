package baijing

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) GetShouyou(req *ghttp.Request) {
	if value, err := g.Redis().DoVar("GET", "BAIJING_SHOUYOU"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.baijingapp.com/index/ajax/getlist/type-3"
	rssData := dao.RSSFeed{
		Title:       "白鲸出海-手游",
		Link:        "https://www.baijingapp.com",
		Description: "白鲸出海手游",
		Tag:         []string{"手游"},
		ImageUrl:    "https://www.baijingapp.com/static/css/default/img/favicon.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := commonHtmlParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", "BAIJING_SHOUYOU", rssStr)
	g.Redis().DoVar("EXPIRE", "BAIJING_SHOUYOU", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
