package guokr

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gurl"
	"regexp"
	"rsshub/app/dao"
	"strconv"
)

type Controller struct {
}

type LinkRouteConfig struct {
	ChannelId string
	Title     string
	LinkType  string
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func commonParser(respString string) (items []dao.RSSItem) {
	respJson := gjson.New(respString)
	dataJsonArray := respJson.Array()
	for index := range dataJsonArray {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string
		dataJson := respJson.GetJson(strconv.Itoa(index))

		title = dataJson.GetString("title")
		link = fmt.Sprintf("https://www.guokr.com/article/%s/", dataJson.GetString("id"))
		author = dataJson.GetString("author.nickname")
		imageLink = dataJson.GetString("small_image")
		content = dataJson.GetString("summary")
		time = dataJson.GetString("date_published")
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

func commonHtmlParser(htmlStr, linkType string) (items []dao.RSSItem) {
	reg := regexp.MustCompile(`window.INITIAL_STORE=(.*?)\n`)
	jsonStr := reg.FindStringSubmatch(htmlStr)
	if len(jsonStr) <= 1 {
		return items
	}
	respJson := gjson.New(jsonStr[1])

	var key string
	if linkType == "calendar" {
		key = "calendarArticleListStore.articleList"
	} else if linkType == "pretty" {
		key = "beautyArticleListStore.articleList"
	}

	dataJsonList := respJson.GetJsons(key)
	for _, dataJson := range dataJsonList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = dataJson.GetString("title")
		link = fmt.Sprintf("https://www.guokr.com/article/%s/", dataJson.GetString("id"))
		author = dataJson.GetString("author.nickname")
		imageLink = dataJson.GetString("small_image")
		imageLink, _ = gurl.Decode(imageLink)
		content = dataJson.GetString("summary")
		time = dataJson.GetString("date_published")
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

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"science": {
			ChannelId: "1",
			LinkType:  "category",
			Title:     "科技"},
		"funny": {
			ChannelId: "2",
			LinkType:  "category",
			Title:     "奇趣"},
		"life": {
			ChannelId: "3",
			LinkType:  "category",
			Title:     "生活"},
		"health": {
			ChannelId: "4",
			LinkType:  "category",
			Title:     "健康"},
		"humanities": {
			ChannelId: "5",
			LinkType:  "category",
			Title:     "人文"},
		"nature": {
			ChannelId: "6",
			LinkType:  "category",
			Title:     "自然"},
		"digital": {
			ChannelId: "8",
			LinkType:  "category",
			Title:     "数码"},
		"food": {
			ChannelId: "9",
			LinkType:  "category",
			Title:     "美食"},

		"engineering": {
			ChannelId: "engineering",
			LinkType:  "subject",
			Title:     "工程"},
		"education": {
			ChannelId: "education",
			LinkType:  "subject",
			Title:     "教育"},
		"physics": {
			ChannelId: "physics",
			LinkType:  "subject",
			Title:     "物理"},
		"sex": {
			ChannelId: "sex",
			LinkType:  "subject",
			Title:     "性情"},
		"agronomy": {
			ChannelId: "agronomy",
			LinkType:  "subject",
			Title:     "农学"},
		"psychology": {
			ChannelId: "psychology",
			LinkType:  "subject",
			Title:     "心理"},
		"medicine": {
			ChannelId: "medicine",
			LinkType:  "subject",
			Title:     "医学"},
		"forensic": {
			ChannelId: "forensic",
			LinkType:  "subject",
			Title:     "法证"},
		"society": {
			ChannelId: "society",
			LinkType:  "subject",
			Title:     "社会"},
		"atmosphere": {
			ChannelId: "atmosphere",
			LinkType:  "subject",
			Title:     "大气"},
		"others": {
			ChannelId: "others",
			LinkType:  "subject",
			Title:     "其他"},
		"chemistry": {
			ChannelId: "chemistry",
			LinkType:  "subject",
			Title:     "化学"},
		"earth": {
			ChannelId: "earth",
			LinkType:  "subject",
			Title:     "地学"},
		"communication": {
			ChannelId: "communication",
			LinkType:  "subject",
			Title:     "传播"},
		"environment": {
			ChannelId: "environment",
			LinkType:  "subject",
			Title:     "环境"},
		"diy": {
			ChannelId: "diy",
			LinkType:  "subject",
			Title:     "diy"},
		"astronomy": {
			ChannelId: "astronomy",
			LinkType:  "subject",
			Title:     "天文"},
		"math": {
			ChannelId: "math",
			LinkType:  "subject",
			Title:     "数学"},
		"biology": {
			ChannelId: "biology",
			LinkType:  "subject",
			Title:     "生物"},
		"aerospace": {
			ChannelId: "aerospace",
			LinkType:  "subject",
			Title:     "航空航天"},
		"internet": {
			ChannelId: "internet",
			LinkType:  "subject",
			Title:     "互联网"},

		"calendar": {
			ChannelId: "calendar",
			LinkType:  "index",
			Title:     "物种日历"},
		"pretty": {
			ChannelId: "pretty",
			LinkType:  "index",
			Title:     "美丽也是技术活"},
	}

	return Links
}
