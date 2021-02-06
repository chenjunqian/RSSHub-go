package juesheng

import (
	"github.com/anaskhan96/soup"
	"rsshub/app/dao"
	"rsshub/lib"
)

type Controller struct {
}

type LinkRouteConfig struct {
	ChannelUrl string
	Title      string
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func commonParser(respString string) (items []dao.RSSItem) {
	respDoc := soup.HTMLParse(respString)
	articleList := respDoc.FindAll("li", "class", "news-item-list")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.Find("img", "class", "effect-zoom").Attrs()["alt"]
		imageLink = article.Find("img", "class", "effect-zoom").Attrs()["data-src"]
		content = article.Find("div", "class", "news-item-brief").Text()
		link = article.Find("a", "class", "news-item-cover").Attrs()["href"]

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
			ChannelUrl: "http://www.juesheng.com/",
			Title:      "快讯"},
		"12k": {
			ChannelUrl: "http://news.juesheng.com/",
			Title:      "12K"},
		"e-edu": {
			ChannelUrl: "http://e.juesheng.com/",
			Title:      "互联网+"},
		"zhijiao": {
			ChannelUrl: "http://news.juesheng.com/zhijiao/",
			Title:      "职业教育"},
		"xingqu": {
			ChannelUrl: "http://news.juesheng.com/xingqu/",
			Title:      "素质教育"},
		"xueqian": {
			ChannelUrl: "http://www.juesheng.com/xqjy/",
			Title:      "学前教育"},
	}

	return Links
}
