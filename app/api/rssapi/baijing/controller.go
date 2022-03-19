package baijing

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
)

type controller struct {
}

var BJController = &controller{}

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
		var (
			title     string
			imageLink string
			link      string
			content   string
		)
		imgDoc := articleDoc.Find("img", "class", "attachment-thumbnail")
		title = imgDoc.Attrs()["title"]
		imageLink = "https://www.baijingapp.com" + imgDoc.Attrs()["src"]
		link = articleDoc.Find("a", "class", "article").Attrs()["href"]
		if !strings.Contains(link, "http") {
			link = "https://www.baijingapp.com" + link
		}
		content = parseCommonDetail(link)
		content = feed.GenerateContent(content)

		rssItem := dao.RSSItem{
			Title:     title,
			Link:      link,
			Content:   content,
			Author:    "",
			Created:   "",
			Thumbnail: imageLink,
		}
		rssItems = append(rssItems, rssItem)
	}

	return
}

func parseCommonDetail(detailLink string) (detailData string) {
	var (
		resp string
	)
	if resp = component.GetContent(detailLink); resp != "" {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)
		respString = resp
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("div", "class", "aw-question-detail")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request baijing common article detail failed, link  %s \n", detailLink)
	}

	return
}
