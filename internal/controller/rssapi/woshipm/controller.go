package woshipm

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
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func commonParser(ctx context.Context, respString string) (items []dao.RSSItem) {
	respJson := gjson.New(respString)
	arrayList := respJson.GetJsons("payload")
	for _, article := range arrayList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.Get("title").String()
		imageLink = article.Get("image").String()
		link = article.Get("permalink").String()
		time = article.Get("date").String()
		author = article.Get("author.name").String()
		content = parseCommonDetail(ctx, link)

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

func parseCommonDetail(ctx context.Context, detailLink string) (detailData string) {
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
		articleElem = docs.Find("div", "class", "article--wrapper")
		if articleElem.Pointer == nil {
			return respString
		}
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf(ctx, "Request woshipm index article detail failed, link  %s \n", detailLink)
	}

	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"recommend": {
			ChannelId: "0",
			Title:     "推荐"},
	}

	return Links
}
