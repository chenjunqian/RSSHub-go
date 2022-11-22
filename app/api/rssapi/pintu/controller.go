package pintu

import (
	"context"
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strconv"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
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

func indexParser(ctx context.Context, respString string) (items []dao.RSSItem) {
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
		title = articleJson.Get("title").String()
		imageLink = articleJson.Get("imgUrl").String()
		time = articleJson.Get("createTime").String()
		content = articleJson.Get("abstracts").String()
		id := articleJson.Get("id")
		op := articleJson.Get("op")
		link = fmt.Sprintf("https://www.pintu360.com/a%s.html?%s", id, op)

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

func parseIndexDetail(ctx context.Context, detailLink string) (detailData string) {
	var (
		resp string
	)
	if resp = component.GetContent(ctx, detailLink); resp != "" {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)

		respString = resp
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("div", "class", "contentLeft")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf(ctx, "Request pintu index article detail failed, link  %s \n", detailLink)
	}

	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"recommend": {
			ChannelId: "0",
			Tags:      []string{"其他"},
			Title:     "推荐"},
		"sell": {
			ChannelId: "7",
			Tags:      []string{"电商"},
			Title:     "零售前沿"},
		"tech": {
			ChannelId: "10",
			Tags:      []string{"科技"},
			Title:     "科技智能"},
		"entertainment": {
			ChannelId: "9",
			Tags:      []string{"娱乐"},
			Title:     "泛文娱"},
		"edu": {
			ChannelId: "98",
			Tags:      []string{"教育"},
			Title:     "教育"},
		"health": {
			ChannelId: "70",
			Tags:      []string{"健康"},
			Title:     "大健康"},
		"consume": {
			ChannelId: "8",
			Tags:      []string{"消费"},
			Title:     "新消费"},
		"startup": {
			ChannelId: "72",
			Tags:      []string{"创业"},
			Title:     "创业投资"},
	}

	return Links
}
