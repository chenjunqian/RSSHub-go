package latexstudio

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *controller) GetArticle(req *ghttp.Request) {

	cacheKey := "LATEXSTUDIO_ARTICLE"
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.latexstudio.net/articles/"
	rssData := dao.RSSFeed{
		Title:       "LATEX 工作室",
		Link:        apiUrl,
		Description: "介绍 LaTeX 的知识和资源分享平台",
		ImageUrl:    "https://www.latexstudio.net/favicon.ico",
	}
	if resp := component.GetContent(apiUrl); resp != "" {
		rssItems := articleParser(resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
