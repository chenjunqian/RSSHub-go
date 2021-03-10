package pintu

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"rsshub/app/dao"
	"rsshub/lib"
	"strconv"
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

func indexParser(respString string) (items []dao.RSSItem) {
	respJson := gjson.New(respString)
	arrayList := respJson.Array()
	for index := 0; index < len(arrayList); index++ {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		articleJson := respJson.GetJson(strconv.Itoa(index))
		title = articleJson.GetString("title")
		imageLink = articleJson.GetString("imgUrl")
		time = articleJson.GetString("createTime")
		content = articleJson.GetString("abstracts")
		id := articleJson.GetString("id")
		op := articleJson.GetString("op")
		link = fmt.Sprintf("https://www.pintu360.com/a%s.html?%s", id, op)

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
		"recommend": {
			ChannelId: "0",
			Title:     "推荐"},
		"sell": {
			ChannelId: "7",
			Title:     "零售前沿"},
		"tech": {
			ChannelId: "10",
			Title:     "科技智能"},
		"entertainment": {
			ChannelId: "9",
			Title:     "泛文娱"},
		"edu": {
			ChannelId: "98",
			Title:     "教育"},
		"health": {
			ChannelId: "70",
			Title:     "大健康"},
		"consume": {
			ChannelId: "8",
			Title:     "新消费"},
		"startup": {
			ChannelId: "72",
			Title:     "创业投资"},
	}

	return Links
}
