package zhihu

import (
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"time"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetAnswers(req *ghttp.Request) {
	peopleId := req.Get("id")
	answerGetUrl := fmt.Sprintf("https://api.zhihu.com/people/%s/answers?order_by=created&offset=0&limit=10", peopleId)
	headers := getHeaders()
	headers["Referer"] = fmt.Sprintf("https://www.zhihu.com/people/%s/answers", peopleId)

	if resp, err := component.GetHttpClient().SetHeaderMap(headers).Get(answerGetUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.GetArray("data")

		// 获取用户名
		peopleName := jsonResp.GetString("data.0.author.name")
		if peopleName == "知乎用户" {
			activitiesUrl := fmt.Sprintf("https://www.zhihu.com/api/v4/members/%s/activities?limit=1", peopleId)
			if activitiesResp, err := component.GetHttpClient().SetHeaderMap(headers).Get(activitiesUrl); err == nil {
				activitiesJsonResp := gjson.New(activitiesResp.ReadAllString())
				peopleName = activitiesJsonResp.GetString("data.0.actor.name")
			}
		}

		rssData := dao.RSSFeed{}
		rssData.Title = fmt.Sprintf("%s的知乎回答", peopleName)
		rssData.Link = fmt.Sprintf("https://www.zhihu.com/people/%s/answers", peopleId)
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"

		items := make([]dao.RSSItem, 0)
		for index := range respDataList {
			title := jsonResp.GetString(fmt.Sprintf("data.%d.question.title", index))
			questionId := jsonResp.GetString(fmt.Sprintf("data.%d.question.id", index))
			questionAnswerId := jsonResp.GetString(fmt.Sprintf("data.%d.id", index))
			link := fmt.Sprintf("https://www.zhihu.com/question/%s/answer/%s", questionId, questionAnswerId)

			// 获取回答内容
			var description string
			answerDetailUrl := fmt.Sprintf("https://api.zhihu.com/appview/api/v4/answers/%s?include=content&is_appview=true", questionAnswerId)
			if answerResp, err := component.GetHttpClient().SetHeaderMap(headers).Get(answerDetailUrl); err == nil {
				answerJsonResp := gjson.New(answerResp.ReadAllString())
				description = answerJsonResp.GetString("content")
				isError := answerJsonResp.GetString("error")
				if isError != "" {
					description = fmt.Sprintf("<a href='%s' target='_blank'>%s</a>", link, title)
				}
			} else {
				description = fmt.Sprintf("<a href='%s' target='_blank'>%s</a>", link, title)
			}

			timeStamp := jsonResp.GetInt64(fmt.Sprintf("data.%d.created_time", index))
			rssItem := dao.RSSItem{
				Title:       title,
				Description: description,
				Author:      peopleName,
				Link:        link,
				Created:     time.Unix(timeStamp, 0).String(),
			}
			items = append(items, rssItem)

		}

		rssData.Items = items
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
