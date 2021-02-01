package jinse

import (
	"github.com/gogf/gf/encoding/gjson"
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

func catalogueParser(jsonString string) (items []dao.RSSItem) {
	respJson := gjson.New(jsonString)
	articleList := respJson.GetJsons("list")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.GetString("title")
		content = article.GetString("extra.summary")
		imageLink = article.GetString("extra.thumbnail_pic")
		author = article.GetString("extra.author")
		time = article.GetString("extra.published_at")
		link = article.GetString("extra.topic_url")

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

func livesParser(jsonString string) (items []dao.RSSItem) {
	respJson := gjson.New(jsonString)
	articleList := respJson.GetJsons("list.0.lives")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.GetString("content")
		content = article.GetString("content")
		imageLink = article.GetString("images.0.url")
		time = article.GetString("created_at")
		link = article.GetString("link")

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

func timelineParser(jsonString string) (items []dao.RSSItem) {
	respJson := gjson.New(jsonString)
	articleList := respJson.GetJsons("data.list")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.GetString("title")
		content = article.GetString("extra.summary")
		imageLink = article.GetString("extra.thumbnail_pic")
		time = article.GetString("extra.published_at")
		link = article.GetString("extra.topic_url")
		author = article.GetString("extra.author")

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
		"zhengce": {
			ChannelId: "zhengce",
			Title:     "政策"},
		"fenxishishuo": {
			ChannelId: "fenxishishuo",
			Title:     "行情"},
		"defi": {
			ChannelId: "defi",
			Title:     "DeFi"},
		"kuang": {
			ChannelId: "kuang",
			Title:     "矿业"},
		"industry": {
			ChannelId: "industry",
			Title:     "产业"},
		"IPFS": {
			ChannelId: "IPFS",
			Title:     "IPFS"},
		"tech": {
			ChannelId: "tech",
			Title:     "技术"},
		"baike": {
			ChannelId: "baike",
			Title:     "百科"},
		"capitalmarket": {
			ChannelId: "capitalmarket",
			Title:     "研报"},
	}

	return Links
}
