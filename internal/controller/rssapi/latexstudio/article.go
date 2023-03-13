package latexstudio

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetArticle(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	cacheKey := "LATEXSTUDIO_ARTICLE"
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.latexstudio.net/articles/"
	rssData := dao.RSSFeed{
		Title:       "LATEX 工作室",
		Link:        apiUrl,
		Description: "介绍 LaTeX 的知识和资源分享平台",
		ImageUrl:    "https://www.latexstudio.net/favicon.ico",
	}
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		rssItems := articleParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx, cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
