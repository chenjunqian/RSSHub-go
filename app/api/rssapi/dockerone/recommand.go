package dockerone

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *controller) GetRecommand(req *ghttp.Request) {

	cacheKey := "DOCKERONE_RECOMMAND"
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "http://weekly.dockone.io/is_recommend-1"
	rssData := dao.RSSFeed{
		Title:       "Dockone",
		Link:        apiUrl,
		Description: "DockOne.io,为技术人员提供最专业的Cloud Native交流平台。",
		ImageUrl:    "http://weekly.dockone.io/static/css/default/img/favicon.ico",
	}
	if resp := component.GetContent(apiUrl); resp != "" {
		rssItems := parseRecommand(resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}