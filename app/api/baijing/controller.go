package baijing

import (
	"github.com/anaskhan96/soup"
	"rsshub/app/dao"
	"rsshub/lib"
	"strings"
)

type Controller struct {
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func commonHtmlParser(htmlStr string) (rssItems []dao.RSSItem) {
	docs := soup.HTMLParse(htmlStr)
	articleDocList := docs.FindAll("div", "class", "articleSingle")
	for _, articleDoc := range articleDocList {
		imgDoc := articleDoc.Find("img", "class", "attachment-thumbnail")
		title := imgDoc.Attrs()["title"]
		imageLink := "https://www.baijingapp.com" + imgDoc.Attrs()["src"]
		link := articleDoc.Find("a", "class", "article").Attrs()["href"]
		if !strings.Contains(link, "http") {
			link = "https://www.baijingapp.com" + link
		}
		var content string
		if articleDoc.Find("div", "class", "articleSingle-content").Error == nil {
			content = articleDoc.Find("div", "class", "articleSingle-content").Find("p").Text()
		}
		description := lib.GenerateDescription(imageLink, content)

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Description: description,
			Author:      "",
			Created:     "",
		}
		rssItems = append(rssItems, rssItem)
	}

	return
}
