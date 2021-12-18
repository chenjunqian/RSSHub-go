package gcore

import (
	"regexp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
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
	articleList := docs.FindAll("div", "class", "col-xl-3")

	for _, article := range articleList {
		var title string
		var link string
		var author string
		var content string
		var time string
		var imageLink string

		title = article.Find("h3", "class", "am_card_title").Text()
		link = "https://www.gcores.com" + article.Find("a", "class", "am_card_content").Attrs()["href"]
		if userInfoEle := article.Find("div", "class", "avatar_text"); userInfoEle.Error == nil {
			author = userInfoEle.Find("h3").Text()
		}
		imageStyle := article.Find("div", "class", "original_imgArea").Attrs()["style"]
		reg := regexp.MustCompile(`background-image:url\((.*?)\)`)
		contentArray := reg.FindStringSubmatch(imageStyle)
		if len(contentArray) > 1 {
			imageLink = contentArray[1]
		}
		content = parseCommonDetail(link)

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: feed.GenerateDescription("", content),
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
		articleElem = docs.Find("div", "class", "newsPage_main")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request gcore article detail failed, link  %s \n", detailLink)
	}

	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"news": {
			ChannelId: "news",
			Tags:      []string{"游戏"},
			Title:     "资讯"},
		"radios": {
			ChannelId: "radios",
			Tags:      []string{"游戏"},
			Title:     "电台"},
		"articles": {
			ChannelId: "articles",
			Tags:      []string{"游戏"},
			Title:     "文章"},
		"videos": {
			ChannelId: "videos",
			Tags:      []string{"游戏", "视频"},
			Title:     "视频"},
	}

	return Links
}
