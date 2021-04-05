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
	Tags      []string
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
			Tags:      []string{"财经", "商业"},
			Title:     "商业产业"},
		"caijing": {
			ChannelId: "caijing",
			Tags:      []string{"财经"},
			Title:     "财经"},
		"dichan": {
			ChannelId: "dichan",
			Tags:      []string{"地产"},
			Title:     "地产"},
		"qiche": {
			ChannelId: "qiche",
			Tags:      []string{"汽车"},
			Title:     "汽车"},
		"tmt": {
			ChannelId: "tmt",
			Tags:      []string{"科技"},
			Title:     "TMT"},
		"guanchajia": {
			ChannelId: "gcj/guanchajia/",
			Tags:      []string{"其他"},
			Title:     "观察家"},
		"zhuanlan": {
			ChannelId: "gcj/zhuanlan/",
			Tags:      []string{"其他"},
			Title:     "专栏"},
		"lishi": {
			ChannelId: "gcj/lishi/",
			Tags:      []string{"历史"},
			Title:     "历史"},
		"shuping": {
			ChannelId: "gcj/shuping/",
			Tags:      []string{"其他"},
			Title:     "书评"},
		"zongshen": {
			ChannelId: "gcj/zongshen/",
			Tags:      []string{"其他"},
			Title:     "纵深"},
		"wenhua": {
			ChannelId: "gcj/wenhua/",
			Tags:      []string{"文化"},
			Title:     "文化"},
		"lingdu": {
			ChannelId: "gcj/lingdu/",
			Tags:      []string{"文化"},
			Title:     "领读"},
	}

	return Links
}
