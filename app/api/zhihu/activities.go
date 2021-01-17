package zhihu

import (
	"fmt"
	"rsshub/app/dao"
	"rsshub/lib"
	"time"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

func (ctl *Controller) GetActivities(req *ghttp.Request) {
	id := req.Get("id")
	url := fmt.Sprintf("https://www.zhihu.com/api/v4/members/%s/activities?limit=7", id)
	headers := getHeaders()
	headers["Referer"] = fmt.Sprintf("https://www.zhihu.com/people/%s/activities", id)
	if resp, err := g.Client().SetHeaderMap(headers).Get(url); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.GetArray("data")

		rssData := dao.RSSFeed{}
		mainTitle := jsonResp.GetString("data.0.actor.name")
		mainDescription := jsonResp.GetString("data.0.actor.headline")
		if mainDescription != "" {
			mainDescription = jsonResp.GetString("data.0.actor.description")
		}
		rssData.Title = fmt.Sprintf("%s的知乎动态", mainTitle)
		rssData.Description = mainDescription
		rssData.Link = url
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"

		items := make([]dao.RSSItem, 0)
		for index := range respDataList {
			var title string
			var description string
			var itemUrl string
			var author string
			contentType := jsonResp.GetString(fmt.Sprintf("data.%d.target.type", index))
			switch contentType {
			case "answer":
				title = jsonResp.GetString(fmt.Sprintf("data.%d.target.question.title", index))
				author = jsonResp.GetString(fmt.Sprintf("data.%d.target.author.name", index))
				// TODO format the content with images
				description = jsonResp.GetString(fmt.Sprintf("data.%d.target.content", index))
				questionId := jsonResp.GetString(fmt.Sprintf("data.%d.target.question.id", index))
				answerId := jsonResp.GetString(fmt.Sprintf("data.%d.target.id", index))
				itemUrl = fmt.Sprintf("https://www.zhihu.com/question/%s/answer/%s", questionId, answerId)
			case "article":
				title = jsonResp.GetString(fmt.Sprintf("data.%d.target.title", index))
				author = jsonResp.GetString(fmt.Sprintf("data.%d.target.author.name", index))
				// TODO format the content with images
				description = jsonResp.GetString(fmt.Sprintf("data.%d.target.content", index))
				questionId := jsonResp.GetString(fmt.Sprintf("data.%d.target.question.id", index))
				answerId := jsonResp.GetString(fmt.Sprintf("data.%d.target.id", index))
				itemUrl = fmt.Sprintf("https://www.zhihu.com/question/%s/answer/%s", questionId, answerId)

			case "pin":
				title = jsonResp.GetString(fmt.Sprintf("data.%d.target.excerpt_title", index))
				author = jsonResp.GetString(fmt.Sprintf("data.%d.target.author.name", index))
				contents := jsonResp.GetArray(fmt.Sprintf("data.%d.target.author.name", index))
				var text string
				var link string
				var images []string
				for contentIndex := range contents {
					pinType := jsonResp.GetString(fmt.Sprintf("data.%d.target.content.%d.type", index, contentIndex))
					switch pinType {
					case "test":
						text = fmt.Sprintf("<p>%s</p>", jsonResp.GetString(fmt.Sprintf("data.%d.target.content.%d.own_text", index, contentIndex)))
					case "image":
						imageItem := jsonResp.GetString(fmt.Sprintf("data.%d.target.content.%d.own_text", index, contentIndex))
						imageHtml := fmt.Sprintf("<p><img src='%s'/></p>", gstr.Replace(imageItem, "xl", "r"))
						images = append(images, imageHtml)
					case "link":
						pinLinkUrl := jsonResp.GetString(fmt.Sprintf("data.%d.target.content.%d.url", index, contentIndex))
						pinLinkTitle := jsonResp.GetString(fmt.Sprintf("data.%d.target.content.%d.title", index, contentIndex))
						link = fmt.Sprintf("<p><a href='%s' target='_blank'>%s</a></p>", pinLinkUrl, pinLinkTitle)
					case "video":
						width := jsonResp.GetString(fmt.Sprintf("data.%d.target.content.%d.playlist.1.width", index, contentIndex))
						height := jsonResp.GetString(fmt.Sprintf("data.%d.target.content.%d.playlist.1.height", index, contentIndex))
						videoUrl := jsonResp.GetString(fmt.Sprintf("data.%d.target.content.%d.playlist.1.url", index, contentIndex))
						link = fmt.Sprintf("<p><video controls='controls' width='%s' height='%s' src='%s'></video></p>", width, height, videoUrl)
					}
				}
				description = fmt.Sprintf("%s%s%s", text, link, gstr.JoinAny(images, ""))
				itemUrl = fmt.Sprintf("https://www.zhihu.com/pin/%s", jsonResp.GetString(fmt.Sprintf("data.%d.target.id", index)))
			case "question":
				title = jsonResp.GetString(fmt.Sprintf("data.%d.target.title", index))
				author = jsonResp.GetString(fmt.Sprintf("data.%d.target.author.name", index))
				// TODO format the content with images
				description = jsonResp.GetString(fmt.Sprintf("data.%d.target.content", index))
				questionId := jsonResp.GetString(fmt.Sprintf("data.%d.target.id", index))
				itemUrl = fmt.Sprintf("https://www.zhihu.com/question/%s", questionId)
			case "collection":
				title = jsonResp.GetString(fmt.Sprintf("data.%d.target.title", index))
				questionId := jsonResp.GetString(fmt.Sprintf("data.%d.target.id", index))
				itemUrl = fmt.Sprintf("https://www.zhihu.com/question/%s", questionId)
			case "column":
				title = jsonResp.GetString(fmt.Sprintf("data.%d.target.title", index))
				// TODO format the content with images
				intro := jsonResp.GetString(fmt.Sprintf("data.%d.target.intro", index))
				imageUrl := jsonResp.GetString(fmt.Sprintf("data.%d.target.image_url", index))
				description = fmt.Sprintf("<p>%s</p><p><img src='%s'/></p>", intro, imageUrl)
				itemUrl = fmt.Sprintf("https://zhuanlan.zhihu.com/%s", jsonResp.GetString(fmt.Sprintf("data.%d.target.id", index)))
			case "topic":
				title = jsonResp.GetString(fmt.Sprintf("data.%d.target.name", index))
				introduction := jsonResp.GetString(fmt.Sprintf("data.%d.target.introduction", index))
				followersCount := jsonResp.GetString(fmt.Sprintf("data.%d.target.followers_count", index))
				description = fmt.Sprintf("<p>%s</p><p>话题关注者人数：%s</p>", introduction, followersCount)
				itemUrl = fmt.Sprintf("https://www.zhihu.com/topic/%s", jsonResp.GetString(fmt.Sprintf("data.%d.target.id", index)))
			case "live":
				title = jsonResp.GetString(fmt.Sprintf("data.%d.target.name", index))
				description = jsonResp.GetString(fmt.Sprintf("data.%d.target.description", index))
				description = gstr.Replace(description, "/\n|\r/g", "<br>")
				itemUrl = fmt.Sprintf("https://www.zhihu.com/lives/%s", jsonResp.GetString(fmt.Sprintf("data.%d.target.id", index)))
			case "roundtable":
				title = jsonResp.GetString(fmt.Sprintf("data.%d.target.name", index))
				description = jsonResp.GetString(fmt.Sprintf("data.%d.target.description", index))
				itemUrl = fmt.Sprintf("https://www.zhihu.com/roundtable/%s", jsonResp.GetString(fmt.Sprintf("data.%d.target.id", index)))

			}

			timeStamp := jsonResp.GetInt64(fmt.Sprintf("data.%d.created_time", index))
			rssItem := dao.RSSItem{
				Title:       title,
				Description: description,
				Author:      author,
				Link:        itemUrl,
				Created:     time.Unix(timeStamp, 0).String(),
			}
			items = append(items, rssItem)
		}

		rssData.Items = items
		rssStr := lib.GenerateRSS(rssData)
		_ = req.Response.WriteXmlExit(rssStr)

	} else {
		_ = req.Response.WriteXmlExit("")
	}

}
