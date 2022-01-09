package guokr

import (
	"fmt"
	"regexp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strconv"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gurl"
	"github.com/gogf/gf/frame/g"
)

type Controller struct {
}

type LinkRouteConfig struct {
	ChannelId string
	Title     string
	LinkType  string
	Tags      []string
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
		content = parseCommonDetail(link)
		time = dataJson.GetString("date_published")
		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: feed.GenerateDescription(content),
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
		articleElem = docs.Find("div", "class", "styled__ArticleContent-sc-1ctyfcr-4")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request guokr article detail failed, link  %s \n", detailLink)
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
		content = parseCommonDetail(link)
		time = dataJson.GetString("date_published")
		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: feed.GenerateDescription(content),
			Created:     time,
			Thumbnail:   imageLink,
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
			Tags:      []string{"科技"},
			Title:     "科技"},
		"funny": {
			ChannelId: "2",
			LinkType:  "category",
			Tags:      []string{"其他"},
			Title:     "奇趣"},
		"life": {
			ChannelId: "3",
			LinkType:  "category",
			Tags:      []string{"生活"},
			Title:     "生活"},
		"health": {
			ChannelId: "4",
			LinkType:  "category",
			Tags:      []string{"健康"},
			Title:     "健康"},
		"humanities": {
			ChannelId: "5",
			LinkType:  "category",
			Tags:      []string{"人文"},
			Title:     "人文"},
		"nature": {
			ChannelId: "6",
			LinkType:  "category",
			Tags:      []string{"其他"},
			Title:     "自然"},
		"digital": {
			ChannelId: "8",
			LinkType:  "category",
			Tags:      []string{"科技"},
			Title:     "数码"},
		"food": {
			ChannelId: "9",
			LinkType:  "category",
			Tags:      []string{"美食"},
			Title:     "美食"},

		"engineering": {
			ChannelId: "engineering",
			LinkType:  "subject",
			Tags:      []string{"科技"},
			Title:     "工程"},
		"education": {
			ChannelId: "education",
			LinkType:  "subject",
			Tags:      []string{"教育"},
			Title:     "教育"},
		"physics": {
			ChannelId: "physics",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "物理"},
		"sex": {
			ChannelId: "sex",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "性情"},
		"agronomy": {
			ChannelId: "agronomy",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "农学"},
		"psychology": {
			ChannelId: "psychology",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "心理"},
		"medicine": {
			ChannelId: "medicine",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "医学"},
		"forensic": {
			ChannelId: "forensic",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "法证"},
		"society": {
			ChannelId: "society",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "社会"},
		"atmosphere": {
			ChannelId: "atmosphere",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "大气"},
		"others": {
			ChannelId: "others",
			LinkType:  "subject",
			Tags:      []string{"其他"},
			Title:     "其他"},
		"chemistry": {
			ChannelId: "chemistry",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "化学"},
		"earth": {
			ChannelId: "earth",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "地学"},
		"communication": {
			ChannelId: "communication",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "传播"},
		"environment": {
			ChannelId: "environment",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "环境"},
		"diy": {
			ChannelId: "diy",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "diy"},
		"astronomy": {
			ChannelId: "astronomy",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "天文"},
		"math": {
			ChannelId: "math",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "数学"},
		"biology": {
			ChannelId: "biology",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "生物"},
		"aerospace": {
			ChannelId: "aerospace",
			LinkType:  "subject",
			Tags:      []string{"科普"},
			Title:     "航空航天"},
		"internet": {
			ChannelId: "internet",
			LinkType:  "subject",
			Tags:      []string{"互联网"},
			Title:     "互联网"},

		"calendar": {
			ChannelId: "calendar",
			LinkType:  "index",
			Tags:      []string{"科普"},
			Title:     "物种日历"},
		"pretty": {
			ChannelId: "pretty",
			LinkType:  "index",
			Tags:      []string{"科普"},
			Title:     "美丽也是技术活"},
	}

	return Links
}
