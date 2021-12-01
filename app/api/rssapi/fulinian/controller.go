package fulinian

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
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
	articleList := docs.FindAll("article", "class", "excerpt")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.Find("img", "class", "thumb").Attrs()["alt"]
		imageLink = article.Find("img", "class", "thumb").Attrs()["src"]
		link = article.Find("a", "class", "focus").Attrs()["href"]
		if category := article.Find("a", "class", "cat"); category.Error == nil {
			author = article.Find("a", "class", "cat").Text()
		}
		time = article.Find("time").Text()

		content = parseCommonDetail(link)

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: feed.GenerateDescription(imageLink, content),
			Created:     time,
		}
		items = append(items, rssItem)
	}
	return
}

func parseCommonDetail(detailLink string) (detailData string) {
	var (
		resp *ghttp.ClientResponse
		err  error
	)
	if resp, err = g.Client().SetHeaderMap(getHeaders()).Get(detailLink); err == nil {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)
		respString = resp.ReadAllString()
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("article", "class", "article-content")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request fulinian article detail failed, link  %s \nerror : %s", detailLink, err)
	}

	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"index": {
			ChannelId: "",
			Tags:      []string{"其他"},
			Title:     "最新发布"},
		"technical-course": {
			ChannelId: "technical-course",
			Tags:      []string{"科技"},
			Title:     "技术教程"},
		"learning": {
			ChannelId: "learning",
			Tags:      []string{"其他"},
			Title:     "学习资料"},
		"chuangye": {
			ChannelId: "chuangye",
			Tags:      []string{"创业"},
			Title:     "创业"},
		"fulinian": {
			ChannelId: "fulinian",
			Tags:      []string{"其他"},
			Title:     "福利年惠"},
		"network-resource": {
			ChannelId: "network-resource",
			Tags:      []string{"科技"},
			Title:     "网络资源"},
		"quality-software": {
			ChannelId: "quality-software",
			Tags:      []string{"互联网", "IT"},
			Title:     "精品软件"},
	}

	return Links
}
