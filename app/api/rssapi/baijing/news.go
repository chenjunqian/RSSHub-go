package baijing

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *controller) GetNews(req *ghttp.Request) {
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
	if resp := component.GetContent(apiUrl); resp != "" {
		rssItems := commonHtmlParser(resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "BAIJING_NEWS", rssStr)
	g.Redis().DoVar("EXPIRE", "BAIJING_NEWS", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
