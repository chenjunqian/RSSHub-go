package zhihu

import (
	"fmt"
	"regexp"
	"rsshub/app/dao"
	"time"

	"rsshub/app/service"

	"github.com/gogf/gf/encoding/gjson"
)

type Controller struct {
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["authority"] = "www.zhihu.com"
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	headers["Authorization"] = "oauth c3cef7c66a1843f8b3a9e6a1e3160e20"
	return headers
}

func getCookieMap() map[string]string {
	cookieMap := service.GetSiteCookies("zhihu")
	return cookieMap
}

func getPinRSSItems(data string) []dao.RSSItem {
	jsonResp := gjson.New(data)
	respDataList := jsonResp.GetArray("data")

	items := make([]dao.RSSItem, 0)
	for index := range respDataList {
		feedItem := dao.RSSItem{}
		targetJson := jsonResp.GetJson(fmt.Sprintf("data.%d.target", index))
		if targetJson == nil {
			targetJson = jsonResp.GetJson(fmt.Sprintf("data.%d", index))
		}
		feedItem.Created = time.Unix(targetJson.GetInt64("created"), 0).String()
		feedItem.Author = targetJson.GetString("author.name")
		feedItem.Title = fmt.Sprintf("%s: %s", feedItem.Author, targetJson.GetString("excerpt_title"))
		feedItem.Link = fmt.Sprintf("https://www.zhihu.com/pin/%s", targetJson.GetString("id"))

		var description string
		contents := targetJson.GetJsons("content")
		for _, content := range contents {
			contentType := content.GetString("type")
			switch contentType {
			case "text":
				description = fmt.Sprintf("%s<div>%s</div>", description, content.GetString("content"))
			case "image":
				contentUrl := content.GetString("url")
				reg := regexp.MustCompile(`_.+\.jpg`)
				contentUrl = reg.ReplaceAllString(contentUrl, `.jpg`)
				description = fmt.Sprintf("%s<img src='%s' />", description, contentUrl)
			case "video":
				width := content.GetString("playlist.hd.width")
				height := content.GetString("playlist.hd.height")
				thumbnail := content.GetString("cover_info.thumbnail")
				playUrl := content.GetString("playlist.hd.play_url")
				description = fmt.Sprintf("%s<video controls='controls' width='%s' height='%s' poster='%s' src='%s'", description, width, height, thumbnail, playUrl)
			case "link":
				linkUrl := content.GetString("url")
				linkTitle := content.GetString("title")
				imageUrl := content.GetString("image_url")
				description = fmt.Sprintf("%s<div><a href='%s'>%s</a><br><img src='%s' /></div>", description, linkUrl, linkTitle, imageUrl)
			}
		}

		endTag := fmt.Sprintf("<a href='https://www.zhihu.com%s'>%s</a>", targetJson.GetString("author.url"), feedItem.Author)
		feedItem.Description = fmt.Sprintf("%s%s", description, endTag)
		items = append(items, feedItem)
	}

	return items
}
