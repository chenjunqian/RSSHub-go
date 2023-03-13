package zhihu

import (
	"context"
	"fmt"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"
	"time"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

func (ctl *Controller) GetActivities(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	id := req.Get("id")
	url := fmt.Sprintf("https://www.zhihu.com/api/v4/members/%s/activities?limit=7", id)
	headers := getHeaders()
	headers["Referer"] = fmt.Sprintf("https://www.zhihu.com/people/%s/activities", id)
	if resp, err := service.GetHttpClient().SetHeaderMap(headers).Get(ctx, url); err == nil {
		defer func(resp *gclient.Response) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}(resp)
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.Get("data").Strings()

		rssData := dao.RSSFeed{}
		mainTitle := jsonResp.Get("data.0.actor.name").String()
		mainDescription := jsonResp.Get("data.0.actor.headline").String()
		if mainDescription != "" {
			mainDescription = jsonResp.Get("data.0.actor.description").String()
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
			contentType := jsonResp.Get(fmt.Sprintf("data.%d.target.type", index)).String()
			switch contentType {
			case "answer":
				title = jsonResp.Get(fmt.Sprintf("data.%d.target.question.title", index)).String()
				author = jsonResp.Get(fmt.Sprintf("data.%d.target.author.name", index)).String()
				// TODO format the content with images
				description = jsonResp.Get(fmt.Sprintf("data.%d.target.content", index)).String()
				questionId := jsonResp.Get(fmt.Sprintf("data.%d.target.question.id", index)).String()
				answerId := jsonResp.Get(fmt.Sprintf("data.%d.target.id", index))
				itemUrl = fmt.Sprintf("https://www.zhihu.com/question/%s/answer/%s", questionId, answerId)
			case "article":
				title = jsonResp.Get(fmt.Sprintf("data.%d.target.title", index)).String()
				author = jsonResp.Get(fmt.Sprintf("data.%d.target.author.name", index)).String()
				// TODO format the content with images
				description = jsonResp.Get(fmt.Sprintf("data.%d.target.content", index)).String()
				questionId := jsonResp.Get(fmt.Sprintf("data.%d.target.question.id", index))
				answerId := jsonResp.Get(fmt.Sprintf("data.%d.target.id", index))
				itemUrl = fmt.Sprintf("https://www.zhihu.com/question/%s/answer/%s", questionId, answerId)

			case "pin":
				title = jsonResp.Get(fmt.Sprintf("data.%d.target.excerpt_title", index)).String()
				author = jsonResp.Get(fmt.Sprintf("data.%d.target.author.name", index)).String()
				contents := jsonResp.Get(fmt.Sprintf("data.%d.target.author.name", index)).Strings()
				var text string
				var link string
				var images []string
				for contentIndex := range contents {
					pinType := jsonResp.Get(fmt.Sprintf("data.%d.target.content.%d.type", index, contentIndex)).String()
					switch pinType {
					case "test":
						text = fmt.Sprintf("<p>%s</p>", jsonResp.Get(fmt.Sprintf("data.%d.target.content.%d.own_text", index, contentIndex)))
					case "image":
						imageItem := jsonResp.Get(fmt.Sprintf("data.%d.target.content.%d.own_text", index, contentIndex)).String()
						imageHtml := fmt.Sprintf("<p><img src='%s'/></p>", gstr.Replace(imageItem, "xl", "r"))
						images = append(images, imageHtml)
					case "link":
						pinLinkUrl := jsonResp.Get(fmt.Sprintf("data.%d.target.content.%d.url", index, contentIndex))
						pinLinkTitle := jsonResp.Get(fmt.Sprintf("data.%d.target.content.%d.title", index, contentIndex))
						link = fmt.Sprintf("<p><a href='%s' target='_blank'>%s</a></p>", pinLinkUrl, pinLinkTitle)
					case "video":
						width := jsonResp.Get(fmt.Sprintf("data.%d.target.content.%d.playlist.1.width", index, contentIndex))
						height := jsonResp.Get(fmt.Sprintf("data.%d.target.content.%d.playlist.1.height", index, contentIndex))
						videoUrl := jsonResp.Get(fmt.Sprintf("data.%d.target.content.%d.playlist.1.url", index, contentIndex))
						link = fmt.Sprintf("<p><video controls='controls' width='%s' height='%s' src='%s'></video></p>", width, height, videoUrl)
					}
				}
				description = fmt.Sprintf("%s%s%s", text, link, gstr.JoinAny(images, ""))
				itemUrl = fmt.Sprintf("https://www.zhihu.com/pin/%s", jsonResp.Get(fmt.Sprintf("data.%d.target.id", index)))
			case "question":
				title = jsonResp.Get(fmt.Sprintf("data.%d.target.title", index)).String()
				author = jsonResp.Get(fmt.Sprintf("data.%d.target.author.name", index)).String()
				// TODO format the content with images
				description = jsonResp.Get(fmt.Sprintf("data.%d.target.content", index)).String()
				questionId := jsonResp.Get(fmt.Sprintf("data.%d.target.id", index))
				itemUrl = fmt.Sprintf("https://www.zhihu.com/question/%s", questionId)
			case "collection":
				title = jsonResp.Get(fmt.Sprintf("data.%d.target.title", index)).String()
				questionId := jsonResp.Get(fmt.Sprintf("data.%d.target.id", index))
				itemUrl = fmt.Sprintf("https://www.zhihu.com/question/%s", questionId)
			case "column":
				title = jsonResp.Get(fmt.Sprintf("data.%d.target.title", index)).String()
				// TODO format the content with images
				intro := jsonResp.Get(fmt.Sprintf("data.%d.target.intro", index))
				imageUrl := jsonResp.Get(fmt.Sprintf("data.%d.target.image_url", index))
				description = fmt.Sprintf("<p>%s</p><p><img src='%s'/></p>", intro, imageUrl)
				itemUrl = fmt.Sprintf("https://zhuanlan.zhihu.com/%s", jsonResp.Get(fmt.Sprintf("data.%d.target.id", index)))
			case "topic":
				title = jsonResp.Get(fmt.Sprintf("data.%d.target.name", index)).String()
				introduction := jsonResp.Get(fmt.Sprintf("data.%d.target.introduction", index))
				followersCount := jsonResp.Get(fmt.Sprintf("data.%d.target.followers_count", index))
				description = fmt.Sprintf("<p>%s</p><p>话题关注者人数：%s</p>", introduction, followersCount)
				itemUrl = fmt.Sprintf("https://www.zhihu.com/topic/%s", jsonResp.Get(fmt.Sprintf("data.%d.target.id", index)))
			case "live":
				title = jsonResp.Get(fmt.Sprintf("data.%d.target.name", index)).String()
				description = jsonResp.Get(fmt.Sprintf("data.%d.target.description", index)).String()
				description = gstr.Replace(description, "/\n|\r/g", "<br>")
				itemUrl = fmt.Sprintf("https://www.zhihu.com/lives/%s", jsonResp.Get(fmt.Sprintf("data.%d.target.id", index)))
			case "roundtable":
				title = jsonResp.Get(fmt.Sprintf("data.%d.target.name", index)).String()
				description = jsonResp.Get(fmt.Sprintf("data.%d.target.description", index)).String()
				itemUrl = fmt.Sprintf("https://www.zhihu.com/roundtable/%s", jsonResp.Get(fmt.Sprintf("data.%d.target.id", index)))

			}

			timeStamp := jsonResp.Get(fmt.Sprintf("data.%d.created_time", index)).Int64()
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
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		req.Response.WriteXmlExit(rssStr)

	} else {
		req.Response.WriteXmlExit("")
	}

}
