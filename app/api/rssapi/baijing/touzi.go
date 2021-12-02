package baijing

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
)

func (ctl *controller) GetTouzi(req *ghttp.Request) {
	if value, err := g.Redis().DoVar("GET", "BAIJING_TOUZI"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.baijingapp.com/index/ajax/getlist/type-10"
	rssData := dao.RSSFeed{
		Title:       "白鲸出海-投融资",
		Link:        "https://www.baijingapp.com",
		Description: "白鲸出海投融资",
		Tag:         []string{"投资"},
		ImageUrl:    "https://www.baijingapp.com/static/css/default/img/favicon.ico",
	}
	if resp, err := component.GetHttpClient().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := commonHtmlParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "BAIJING_TOUZI", rssStr)
	g.Redis().DoVar("EXPIRE", "BAIJING_TOUZI", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
