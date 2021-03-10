package medsci

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

func commonParser(respString string) (items []dao.RSSItem) {
	respDoc := soup.HTMLParse(respString)
	articleList := respDoc.FindAll("div", "class", "item")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		if imgATag := article.Find("a", "class", "ms-statis"); imgATag.Error == nil {
			if imageDiv := imgATag.Find("img"); imageDiv.Error == nil {
				imageLink = imageDiv.Attrs()["src"]
				title = imageDiv.Attrs()["title"]
				link = imgATag.Attrs()["href"]
			}
		}

		if title == "" {
			continue
		}

		if contentEle := article.Find("p", "class", "text-justify"); contentEle.Error == nil {
			content = contentEle.Text()
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
			ChannelId: "http://www.juesheng.com/",
			Title:     "快讯"},
	}

	return Links
}
