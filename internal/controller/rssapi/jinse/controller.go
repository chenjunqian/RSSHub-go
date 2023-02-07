package jinse

import (
	"context"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct {
}

type LinkRouteConfig struct {
	ChannelId string
	Title     string
	Tags      []string
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func catalogueParser(ctx context.Context, jsonString string) (items []dao.RSSItem) {
	respJson := gjson.New(jsonString)
	articleList := respJson.GetJsons("list")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.Get("title").String()
		imageLink = article.Get("extra.thumbnail_pic").String()
		author = article.Get("extra.author").String()
		time = article.Get("extra.published_at").String()
		link = article.Get("extra.topic_url").String()
		content = parseCatalogueDetail(ctx, link)

		rssItem := dao.RSSItem{
			Title:     title,
			Link:      link,
			Author:    author,
			Content:   feed.GenerateContent(content),
			Created:   time,
			Thumbnail: imageLink,
		}
		items = append(items, rssItem)
	}
	return
}

func parseCatalogueDetail(ctx context.Context, detailLink string) (detailData string) {
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
		articleElem = docs.Find("div", "class", "js-article")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf(ctx, "Request jinse Catalogue article detail failed, link  %s \n", detailLink)
	}

	return
}

func livesParser(ctx context.Context, jsonString string) (items []dao.RSSItem) {
	respJson := gjson.New(jsonString)
	articleList := respJson.GetJsons("list.0.lives")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.Get("content").String()
		imageLink = article.Get("images.0.url").String()
		time = article.Get("created_at").String()
		link = article.Get("link").String()
		content = parseLiveDetail(ctx, link)

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: feed.GenerateContent(content),
			Created:     time,
			Thumbnail:   imageLink,
		}
		items = append(items, rssItem)
	}
	return
}

func parseLiveDetail(ctx context.Context, detailLink string) (detailData string) {
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
		articleElem = docs.Find("section", "class", "at-body")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf(ctx, "Request jinse Catalogue article detail failed, link  %s \n", detailLink)
	}

	return
}

func timelineParser(ctx context.Context, jsonString string) (items []dao.RSSItem) {
	respJson := gjson.New(jsonString)
	articleList := respJson.GetJsons("data.list")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.Get("title").String()
		imageLink = article.Get("extra.thumbnail_pic").String()
		time = article.Get("extra.published_at").String()
		link = article.Get("extra.topic_url").String()
		author = article.Get("extra.author").String()
		content = parseTimelineDetail(ctx, link)

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: feed.GenerateContent(content),
			Created:     time,
			Thumbnail:   imageLink,
		}
		items = append(items, rssItem)
	}
	return
}

func parseTimelineDetail(ctx context.Context, detailLink string) (detailData string) {
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
		articleElem = docs.Find("div", "class", "js-article")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf(ctx, "Request jinse Catalogue article detail failed, link  %s \n", detailLink)
	}

	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"zhengce": {
			ChannelId: "zhengce",
			Tags:      []string{"政治"},
			Title:     "政策"},
		"fenxishishuo": {
			ChannelId: "fenxishishuo",
			Tags:      []string{"财经"},
			Title:     "行情"},
		"defi": {
			ChannelId: "defi",
			Tags:      []string{"财经"},
			Title:     "DeFi"},
		"kuang": {
			ChannelId: "kuang",
			Tags:      []string{"财经"},
			Title:     "矿业"},
		"industry": {
			ChannelId: "industry",
			Tags:      []string{"财经"},
			Title:     "产业"},
		"IPFS": {
			ChannelId: "IPFS",
			Tags:      []string{"财经"},
			Title:     "IPFS"},
		"tech": {
			ChannelId: "tech",
			Tags:      []string{"技术"},
			Title:     "技术"},
		"baike": {
			ChannelId: "baike",
			Tags:      []string{"财经"},
			Title:     "百科"},
		"capitalmarket": {
			ChannelId: "capitalmarket",
			Tags:      []string{"财经"},
			Title:     "研报"},
	}

	return Links
}
