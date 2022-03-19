package yanxishe

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

type controller struct {
}

var Controller = &controller{}

func parseIndex(respString string) (items []dao.RSSItem) {
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

			content = parseIndexDetail(link)
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

func parseIndexDetail(link string) (detailData string) {
	respString := component.GetContent(link)
	respDoc := soup.HTMLParse(respString)
	articleDoc := respDoc.Find("div", "class", "lph-article-comView")
	if articleDoc.Pointer != nil {
		detailData = articleDoc.HTML()
	} else {
		g.Log().Error(link)
	}

	return
}
