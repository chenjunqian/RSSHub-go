package guanchajia

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

func commonParser(htmlStr string) (items []dao.RSSItem) {
	docs := soup.HTMLParse(htmlStr)
	articleList := docs.Find("ul", "id", "lyp_article").FindAll("li")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.Find("div").Find("span").Find("a").Text()
		imageLink = article.Find("img").Attrs()["src"]
		link = article.Find("div").Find("span").Find("a").Attrs()["href"]
		content = article.Find("p").Text()

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
		"shangyechanye": {
			ChannelId: "shangyechanye",
			Title:     "商业产业"},
		"caijing": {
			ChannelId: "caijing",
			Title:     "财经"},
		"dichan": {
			ChannelId: "dichan",
			Title:     "地产"},
		"qiche": {
			ChannelId: "qiche",
			Title:     "汽车"},
		"tmt": {
			ChannelId: "tmt",
			Title:     "TMT"},
		"guanchajia": {
			ChannelId: "gcj/guanchajia/",
			Title:     "观察家"},
		"zhuanlan": {
			ChannelId: "gcj/zhuanlan/",
			Title:     "专栏"},
		"lishi": {
			ChannelId: "gcj/lishi/",
			Title:     "历史"},
		"shuping": {
			ChannelId: "gcj/shuping/",
			Title:     "书评"},
		"zongshen": {
			ChannelId: "gcj/zongshen/",
			Title:     "纵深"},
		"wenhua": {
			ChannelId: "gcj/wenhua/",
			Title:     "文化"},
		"lingdu": {
			ChannelId: "gcj/lingdu/",
			Title:     "领读"},
	}

	return Links
}
