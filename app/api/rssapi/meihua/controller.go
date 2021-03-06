package meihua

import (
	"github.com/anaskhan96/soup"
	"regexp"
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

func commonParser(respString string) (items []dao.RSSItem) {
	respDoc := soup.HTMLParse(respString)
	articleList := respDoc.FindAll("div", "class", "article-item")
	baseUrl := "https://www.meihua.info"
	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		coverDiv := article.Find("div", "class", "article-item-cover")
		imgDoc := coverDiv.Find("img")
		title = imgDoc.Attrs()["alt"]
		coverDivString := coverDiv.HTML()
		reg := regexp.MustCompile(`<noscript><img src="(.*?)"`)
		imageUrlMatches := reg.FindStringSubmatch(coverDivString)
		if len(imageUrlMatches) > 1 {
			imageLink = imageUrlMatches[1]
		}

		link = baseUrl + coverDiv.Find("a").Attrs()["href"]

		content = article.Find("div", "class", "intro").Text()
		spanList := article.FindAll("span", "class", "text-tag")
		for _, spanDoc := range spanList {
			if aTag := spanDoc.Find("a"); aTag.Error == nil {
				author = aTag.Text()
			}
			if spanTag := spanDoc.Find("span"); spanTag.Error == nil {
				time = spanTag.Text()
			}
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
		"latest": {
			ChannelId: "article!2",
			Title:     "最新"},
		"hot": {
			ChannelId: "article!3",
			Title:     "热门"},
	}

	return Links
}
