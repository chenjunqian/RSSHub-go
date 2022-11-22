package _199IT

import (
	"context"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/frame/g"
)

type controller struct {
}

var IT1999Controller = &controller{}

func parseArticle(ctx context.Context,htmlStr string) []dao.RSSItem {
	doc := soup.HTMLParse(htmlStr)
	articleList := doc.FindAll("article")
	rssItems := make([]dao.RSSItem, 0)
	for _, article := range articleList {
		var (
			title       string
			thumbnail   string
			description string
			time        string
			link        string
			detailData  string
		)
		title = article.Find("h2", "class", "entry-title").Find("a").Attrs()["title"]
		thumbnail = article.Find("img", "class", "attachment-post-thumbnail").Attrs()["src"]
		time = article.Find("time", "class", "entry-date").Attrs()["datetime"]
		link = article.Find("h2", "class", "entry-title").Find("a").Attrs()["href"]
		detailData = parseDetail(ctx, link)
		description = feed.GenerateContent(detailData)
		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Description: description,
			Content:     detailData,
			Created:     time,
			Thumbnail:   thumbnail,
		}
		rssItems = append(rssItems, rssItem)
	}

	return rssItems
}

func parseDetail(ctx context.Context,detailLink string) (detailData string) {
	var (
		resp string
	)
	if resp = component.GetContent(ctx,detailLink); resp != "" {
		var (
			docs        soup.Root
			articleElem soup.Root
		)
		docs = soup.HTMLParse(resp)
		articleElem = docs.Find("article")
		detailData = articleElem.HTML()
	} else {
		g.Log().Errorf(ctx,"Request 199IT article detail failed, link  %s \n", detailLink)
	}

	return
}
