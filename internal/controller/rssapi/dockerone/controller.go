package dockerone

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/frame/g"
)

type controller struct {
}

var Controller = &controller{}

func parseRecommand(ctx context.Context, respString string) (items []dao.RSSItem) {
	respDoc := soup.HTMLParse(respString)
	articleListWarpper := respDoc.Find("div", "class", "aw-common-list")
	articleList := articleListWarpper.FindAll("div", "class", "aw-item")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		if titleDoc := article.Find("h4"); titleDoc.Pointer != nil {
			title = titleDoc.Find("a").Text()
			link = titleDoc.Find("a").Attrs()["href"]
		}

		author = article.Find("a", "class", "aw-user-name").Text()
		dateDoc := article.FindAll("span", "class", "text-color-999")[0].Text()
		timeArray := strings.Split(dateDoc, " • ")
		time = timeArray[len(timeArray)-1]
		content = parseRecommandDetail(ctx, link)

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

func parseRecommandDetail(ctx context.Context, link string) (detailData string) {
	respString := service.GetContent(ctx, link)
	respDoc := soup.HTMLParse(respString)
	articleDoc := respDoc.Find("div", "class", "markitup-box")
	if articleDoc.Pointer != nil {
		detailData = articleDoc.HTML()
	} else {
		g.Log().Error(ctx, link)
	}

	return
}
