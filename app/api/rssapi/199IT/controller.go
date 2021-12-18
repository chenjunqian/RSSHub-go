package _199IT

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
)

type controller struct {
}

var IT1999Controller = &controller{}

func parseArticle(htmlStr string) []dao.RSSItem {
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
		detailData = parseDetail(link)
		description = feed.GenerateDescription(thumbnail, detailData)
		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Description: description,
			Created:     time,
			Thumbnail:   thumbnail,
		}
		rssItems = append(rssItems, rssItem)
	}

	return rssItems
}

func parseDetail(detailLink string) (detailData string) {
	var (
		resp string
	)
	if resp = component.GetContent(detailLink); resp != "" {
		var (
			docs        soup.Root
			articleElem soup.Root
		)
		docs = soup.HTMLParse(resp)
		articleElem = docs.Find("article")
		detailData = articleElem.HTML()
	} else {
		g.Log().Errorf("Request 199IT article detail failed, link  %s \n", detailLink)
	}

	return
}
