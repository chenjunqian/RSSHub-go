package sspai

import (
	"github.com/gogf/gf/encoding/gjson"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
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
	articleList := respJson.GetJsons("data")
	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.GetString("title")
		content = article.GetString("summary")
		imageLink = "https://cdn.sspai.com/" + article.GetString("banner")
		author = article.GetString("author.nickname")
		time = article.GetString("created_time")
		link = "https://sspai.com/post/" + article.GetString("id")

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: feed.GenerateDescription(imageLink, content),
			Created:     time,
			Thumbnail:   imageLink,
		}
		items = append(items, rssItem)
	}
	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"recommend": {
			ChannelId: "recommend",
			Title:     "推荐"},
		"hot": {
			ChannelId: "hot",
			Title:     "热门文章"},
		"app_recommend": {
			ChannelId: "app_recommend",
			Title:     "应用推荐"},
		"skill": {
			ChannelId: "skill",
			Title:     "效率技巧"},
		"lifestyle": {
			ChannelId: "lifestyle",
			Title:     "生活方式"},
		"podcast": {
			ChannelId: "podcast",
			Title:     "少数派播客"},
	}

	return Links
}
