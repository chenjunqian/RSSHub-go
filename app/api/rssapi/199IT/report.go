package _199IT

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) Get199ITCategoryReport(req *ghttp.Request) {
	if value, err := g.Redis().DoVar("GET", "199IT_CATEGORY_REPORT"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "http://www.199it.com/archives/category/report"
	headers := getHeaders()
	rssData := dao.RSSFeed{
		Title:       "199it",
		Link:        apiUrl,
		Tag:         []string{"互联网", "IT", "科技"},
		ImageUrl:    "http://www.199it.com/favicon.ico",
		Description: "互联网数据资讯网-研究报告-199IT",
	}
	if resp, err := g.Client().SetHeaderMap(headers).Get(apiUrl); err == nil {
		rssItems := parseArticle(resp.ReadAllString())
		rssData.Items = rssItems
	}
	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", "199IT_CATEGORY_REPORT", rssStr)
	g.Redis().DoVar("EXPIRE", "199IT_CATEGORY_REPORT", 60*10)
	_ = req.Response.WriteXmlExit(rssStr)
}
