package oschina

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *controller) GetLatestNews(req *ghttp.Request) {
	var (
		cacheKey string
		apiUrl   string
	)
	cacheKey = "OSCHINA_NEWS_LATEST"
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl = "https://www.oschina.net/news"
	rssData := dao.RSSFeed{
		Title:       "掘金推荐 - 最热",
		Link:        apiUrl,
		Description: "开源,OSC,开源软件,开源硬件,开源网站,开源社区,java开源,perl开源,python开源,ruby开源,php开源,开源项目,开源代码",
		ImageUrl:    "https://static.oschina.net/new-osc/img/favicon.ico",
	}

	if resp := component.GetContentByMobile(apiUrl); resp != "" {
		rssItems := parseNews(resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}

func parseNews(respString string) (items []dao.RSSItem) {
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
		item.Description = feed.GenerateDescription("", parseNewsDetail(item.Link))
		items = append(items, item)
	}
	return
}

func parseNewsDetail(detailLink string) (detailData string) {
	var (
		resp string
	)
	if resp = component.GetContentByMobile(detailLink); resp != "" {
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
		g.Log().Errorf("Request oschina news article detail failed, link  %s \n", detailLink)
	}

	return
}
