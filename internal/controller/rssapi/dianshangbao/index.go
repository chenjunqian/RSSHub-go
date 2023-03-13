package dianshangbao

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetIndex(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getNewsLinks()[linkType]

	cacheKey := "DSB_INDEX_" + linkType
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
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
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		respDocs := soup.HTMLParse(resp)
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

			content = parseNewsDetail(ctx, link)

			rssItem := dao.RSSItem{
				Title:     title,
				Link:      link,
				Content:   feed.GenerateContent(content),
				Thumbnail: imageLink,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx, cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}

func parseNewsDetail(ctx context.Context, detailLink string) (detailData string) {
	var (
		resp string
	)
	if resp = service.GetContent(ctx, detailLink); resp != "" {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)
		respString = resp
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("div", "class", "article-content-container")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf(ctx, "Request dianshangbao news article detail failed, link  %s \n", detailLink)
	}

	return
}
