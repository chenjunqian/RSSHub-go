package baijing

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) GetNews(req *ghttp.Request) {
	if value, err := g.Redis().DoVar("GET", "BAIJING_NEWS"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.baijingapp.com"
	rssData := dao.RSSFeed{
		Title:       "白鲸出海-快讯",
		Link:        apiUrl,
		Description: "白鲸出海快讯",
		Tag:         []string{"新闻"},
		ImageUrl:    "https://www.baijingapp.com/static/css/default/img/favicon.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := commonHtmlParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", "BAIJING_NEWS", rssStr)
	g.Redis().DoVar("EXPIRE", "BAIJING_NEWS", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
