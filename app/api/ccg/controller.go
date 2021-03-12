package ccg

import (
	"github.com/anaskhan96/soup"
	"rsshub/app/dao"
	"rsshub/lib"
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

func indexParser(respString string) (items []dao.RSSItem) {
	respSoup := soup.HTMLParse(respString)
	var articleList []soup.Root
	if articleUl := respSoup.Find("ul", "class", "huodong-list"); articleUl.Error != nil {
		return
	} else {
		articleList = articleUl.FindAll("li")
	}
	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		if titleH5 := article.Find("h5"); titleH5.Error == nil {
			title = titleH5.Text()
		} else {
			continue
		}

		if imageHtml := article.Find("img"); imageHtml.Error == nil {
			imageLink = imageHtml.Attrs()["src"]
		}

		if linkHtml := article.Find("a"); linkHtml.Error == nil {
			link = linkHtml.Attrs()["href"]
		}

		if contentHtml := article.Find("p"); contentHtml.Error == nil {
			content = contentHtml.Text()
		}

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: lib.GenerateDescription(imageLink, content),
			Created:     time,
		}
		items = append(items, rssItem)
	}
	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"news": {
			ChannelId: "news",
			Title:     "新闻动态"},
		"media": {
			ChannelId: "mtbd",
			Title:     "媒体报道"},
	}

	return Links
}
