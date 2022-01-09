package woshipm

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
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
	respJson := gjson.New(respString)
	arrayList := respJson.GetJsons("payload")
	for _, article := range arrayList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.GetString("title")
		imageLink = article.GetString("image")
		link = article.GetString("permalink")
		time = article.GetString("date")
		author = article.GetString("author.name")
		content = parseCommonDetail(link)

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: feed.GenerateDescription(content),
			Created:     time,
			Thumbnail:   imageLink,
		}
		items = append(items, rssItem)
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
		articleElem = docs.Find("div", "class", "article--wrapper")
		if articleElem.Pointer == nil {
			return respString
		}
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request woshipm index article detail failed, link  %s \n", detailLink)
	}

	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"recommend": {
			ChannelId: "0",
			Title:     "推荐"},
	}

	return Links
}
