package duozhi

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"rsshub/app/dao"
	"strings"
)

type Controller struct {
}

type IndustryNewsRouteConfig struct {
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
	respDocs := soup.HTMLParse(htmlStr)
	dataDocsList := respDocs.FindAll("div", "class", "post-item")
	if len(dataDocsList) > 20 {
		dataDocsList = dataDocsList[:20]
	}
	for _, dataDocs := range dataDocsList {
		var imageLink string
		var title string
		var link string
		var content string
		var author string
		var time string

		postImageWrap := dataDocs.Find("a", "class", "post-img")
		if postImageWrap.Error == nil {
			title = postImageWrap.Attrs()["title"]
			link = postImageWrap.Attrs()["href"]
			imageStyleStr := postImageWrap.Attrs()["style"]
			imageStyleStrs := strings.Split(imageStyleStr, "url(")
			if len(imageStyleStrs) >= 2 {
				imageLink = imageStyleStrs[1]
			}
		}
		contentWrap := dataDocs.Find("p", "class", "post-desc")
		if contentWrap.Error == nil {
			content = contentWrap.Text()
		}
		authorWrap := dataDocs.Find("span", "class", "post-attr")
		if authorWrap.Error == nil {
			author = authorWrap.Text()
		}

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: fmt.Sprintf("<img src='%s'><br>%s", imageLink, content),
			Created:     time,
		}
		items = append(items, rssItem)
	}
	return
}

func getIndustryNewsLinks() map[string]IndustryNewsRouteConfig {
	Links := map[string]IndustryNewsRouteConfig{
		"insight": {
			ChannelId: "insight",
			Title:     "观察"},
		"preschool": {
			ChannelId: "preschool",
			Title:     "早教"},
		"K12": {
			ChannelId: "K12",
			Title:     "K12"},
		"qualityedu": {
			ChannelId: "qualityedu",
			Title:     "素质教育"},
		"adultedu": {
			ChannelId: "adult",
			Title:     "职教/大学生"},
		"EduInformatization": {
			ChannelId: "EduInformatization",
			Title:     "信息化教育"},
		"earnings": {
			ChannelId: "earnings",
			Title:     "财报"},
		"privateschools": {
			ChannelId: "privateschools",
			Title:     "民办教育"},
		"overseas": {
			ChannelId: "overseas",
			Title:     "留学"},
	}

	return Links
}
