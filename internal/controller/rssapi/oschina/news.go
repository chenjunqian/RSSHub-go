package oschina

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetLatestNews(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	var (
		cacheKey string
		apiUrl   string
	)
	cacheKey = "OSCHINA_NEWS_LATEST"
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl = "https://www.oschina.net/news"
	rssData := dao.RSSFeed{
		Title:       "掘金推荐 - 最热",
		Link:        apiUrl,
		Description: "开源,OSC,开源软件,开源硬件,开源网站,开源社区,java开源,perl开源,python开源,ruby开源,php开源,开源项目,开源代码",
		ImageUrl:    "https://static.oschina.net/new-osc/img/favicon.ico",
	}

	if resp := service.GetContentByMobile(ctx, apiUrl); resp != "" {
		rssItems := parseNews(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx, cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}

func parseNews(ctx context.Context, respString string) (items []dao.RSSItem) {
	var (
		docs        soup.Root
		articleList []soup.Root
	)
	docs = soup.HTMLParse(respString)
	articleList = docs.FindAll("div", "class", "news-item")
	for _, articleItemDoc := range articleList {
		var item dao.RSSItem
		var headerDoc soup.Root
		headerDoc = articleItemDoc.Find("h3", "class", "header")
		if headerDoc.Pointer != nil {
			var titleDoc = headerDoc.Find("a")
			if titleDoc.Pointer != nil {
				item.Title = titleDoc.Text()
				item.Link = titleDoc.Attrs()["href"]
			}
		}
		item.Content = feed.GenerateContent(parseNewsDetail(ctx, item.Link))
		items = append(items, item)
	}
	return
}

func parseNewsDetail(ctx context.Context, detailLink string) (detailData string) {
	var (
		resp string
	)
	if resp = service.GetContentByMobile(ctx, detailLink); resp != "" {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)

		respString = resp
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("div", "class", "article-detail")
		if articleElem.Pointer != nil {
			detailData = articleElem.HTML()
		}

	} else {
		g.Log().Errorf(ctx, "Request oschina news article detail failed, link  %s \n", detailLink)
	}

	return
}
