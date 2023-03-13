package whalegogo

import (
	"context"
	"fmt"
	"rsshub/internal/dao"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
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

func indexParser(ctx context.Context, jsonString string) (items []dao.RSSItem) {
	respJson := gjson.New(jsonString)
	articleList := respJson.GetJsons("response.items.item")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.Get("title").String()
		content = article.Get("description").String()
		imageLink = article.Get("cover").String()
		author = article.Get("author.username").String()
		time = article.Get("created_at").String()
		id := article.Get("id")
		cId := article.Get("cid")
		link = fmt.Sprintf("https://m.whalegogo.com/article?id=%s&cid=%s", id, cId)

		rssItem := dao.RSSItem{
			Title:     title,
			Link:      link,
			Author:    author,
			Content:   feed.GenerateContent(content),
			Created:   time,
			Thumbnail: imageLink,
		}
		items = append(items, rssItem)
	}
	return
}

func portalParser(jsonString string) (items []dao.RSSItem) {
	respJson := gjson.New(jsonString)
	fmt.Println(jsonString)
	articleList := respJson.GetJsons("response.items.item")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.Get("title").String()
		content = article.Get("description").String()
		imageLink = article.Get("cover").String()
		author = article.Get("author.username").String()
		time = article.Get("created_at").String()
		id := article.Get("id")
		link = fmt.Sprintf("https://www.whalegogo.com/news?id=%s", id)

		rssItem := dao.RSSItem{
			Title:     title,
			Link:      link,
			Author:    author,
			Content:   feed.GenerateContent(content),
			Created:   time,
			Thumbnail: imageLink,
		}
		items = append(items, rssItem)
	}
	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"news": {
			ChannelId: "2",
			Title:     "快讯"},
		"article": {
			ChannelId: "1",
			Title:     "文章"},
		"activities": {
			ChannelId: "7",
			Title:     "活动"},
		"appraisal": {
			ChannelId: "8",
			Title:     "评测"},
	}

	return Links
}
