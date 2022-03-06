package latexstudio

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/anaskhan96/soup"
)

type controller struct {
}

var Controller = &controller{}

func articleParser(respString string) (items []dao.RSSItem) {
	respDoc := soup.HTMLParse(respString)
	articleListWarpper := respDoc.Find("div", "class", "article-list")
	articleList := articleListWarpper.FindAll("div", "class", "media")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		leftDoc := article.Find("div", "class", "media-left")
		bodyDoc := article.Find("div", "class", "media-body")
		imageLink = leftDoc.Find("img").Attrs()["src"]

		if titleDoc := bodyDoc.Find("h3"); titleDoc.Pointer != nil {
			title = titleDoc.Find("a").Text()
			link = titleDoc.Find("a").Attrs()["href"]
			content = articleDetailParser(link)
		}

		if authorDoc := bodyDoc.Find("div", "class", "article-tag"); authorDoc.Pointer != nil {
			author = authorDoc.Find("a").Text()
			time = authorDoc.Find("span", "itemprop", "date").Text()
			time = strings.ReplaceAll(time, "年", "/")
			time = strings.ReplaceAll(time, "月", "/")
			time = strings.ReplaceAll(time, "日", "/")
		}

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: feed.GenerateDescription(content),
			Created:     time,
			Thumbnail:   imageLink,
		}
		items = append(items, rssItem)
	}
	return
}

func articleDetailParser(link string) (detailData string) {
	respString := component.GetContent(link)
	respDoc := soup.HTMLParse(respString)
	articleDoc := respDoc.Find("div", "class", "article-text")
	detailData = articleDoc.HTML()
	return
}
