package dongqiudi

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
)

func (ctl *Controller) GetDaily(req *ghttp.Request) {

	cacheKey := "DONGQIUDI_DAILY"
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.dongqiudi.com/special/48"
	rssData := dao.RSSFeed{
		Title:       "懂球帝 - 早报",
		Link:        apiUrl,
		Tag:         []string{"体育"},
		Description: "早报 — 专题|专业权威的足球网站|懂球帝",
		ImageUrl:    "https://static1.dongqiudi.com/web-new/web/images/fav.ico",
	}
	if resp, err := component.GetHttpClient().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		defer func(resp *ghttp.ClientResponse) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(err)
			}
		}(resp)
		rssData.Items = commonParser(resp.ReadAllString())
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*6)
	_ = req.Response.WriteXmlExit(rssStr)
}
