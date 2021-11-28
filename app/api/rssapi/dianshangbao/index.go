package dianshangbao

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
	"strings"
)

func (ctl *controller) GetIndex(req *ghttp.Request) {

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getNewsLinks()[linkType]

	cacheKey := "DSB_INDEX_" + linkType
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.dsb.cn/" + linkConfig.ChannelId
	rssData := dao.RSSFeed{
		Title:       "电商报 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         linkConfig.Tags,
		Description: "电商报行业观察栏目，重点针对电子商务行业、互联网行业、it行业重大新闻24小时跟踪报道，揭示电子商务、互联网等行业的发展趋势和分析报告。",
		ImageUrl:    "https://www.dsb.cn/favicon.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		respDocs := soup.HTMLParse(resp.ReadAllString())
		dataDocsList := respDocs.FindAll("li", "class", "clearfix")
		rssItems := make([]dao.RSSItem, 0)
		for _, dataDocs := range dataDocsList {
			var imageLink, title, link, content string
			imgDoc := dataDocs.Find("img")
			if imgDoc.Error == nil {
				imageLink = dataDocs.Find("img").Attrs()["src"]
				title = dataDocs.Find("img").Attrs()["title"]
			} else {
				continue
			}
			contentDoc := dataDocs.Find("div", "class", "new-list-con-r")
			if contentDoc.Error == nil && contentDoc.Find("a").Error == nil {
				link = contentDoc.Find("a").Attrs()["href"]
			} else {
				continue
			}

			content = parseNewsDetail(link)

			rssItem := dao.RSSItem{
				Title:       title,
				Link:        link,
				Description: lib.GenerateDescription(imageLink, content),
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := lib.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}

func parseNewsDetail(detailLink string) (detailData string) {
	var (
		resp *ghttp.ClientResponse
		err  error
	)
	if resp, err = g.Client().SetHeaderMap(getHeaders()).Get(detailLink); err == nil {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)
		respString = resp.ReadAllString()
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("div", "class", "article-content-container")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request dianshangbao news article detail failed, link  %s \nerror : %s", detailLink, err)
	}

	return
}
