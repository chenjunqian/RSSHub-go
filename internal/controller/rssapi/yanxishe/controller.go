package yanxishe

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type controller struct {
}

var Controller = &controller{}

func parseIndex(ctx context.Context, respString string) (items []dao.RSSItem) {
	respDoc := soup.HTMLParse(respString)
	articleListWarpper := respDoc.Find("div", "class", "lph-pageList")
	articleList := articleListWarpper.FindAll("div", "class", "box")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		if imgDoc := article.Find("div", "class", "img"); imgDoc.Pointer != nil {
			imageLink = imgDoc.Find("img", "class", "lazy").Attrs()["src"]
			title = imgDoc.Find("img", "class", "lazy").Attrs()["title"]
			aTags := imgDoc.FindAll("a")
			if len(aTags) > 1 {
				link = aTags[1].Attrs()["href"]
			}

			content = parseIndexDetail(ctx, link)
		}

		if wordDoc := article.Find("div", "class", "word"); wordDoc.Pointer != nil {
			author = wordDoc.Find("a", "class", "aut").Text()
			time = wordDoc.Find("div", "class", "time").Text()
			time = strings.ReplaceAll(time, "月", "/")
			time = strings.ReplaceAll(time, "日", "")
			time = strconv.Itoa(gtime.Now().Year()) + "/" + time
		}

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

func parseIndexDetail(ctx context.Context, link string) (detailData string) {
	respString := service.GetContent(ctx, link)
	respDoc := soup.HTMLParse(respString)
	articleDoc := respDoc.Find("div", "class", "lph-article-comView")
	if articleDoc.Pointer != nil {
		detailData = articleDoc.HTML()
	} else {
		g.Log().Error(ctx, link)
	}

	return
}
