package zhihu

import (
	"fmt"
	"rsshub/lib"
	"time"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetAnswers(req *ghttp.Request) {
	peopleId := req.Get("id")
	answerGetUrl := fmt.Sprintf("https://api.zhihu.com/people/%s/answers?order_by=created&offset=0&limit=10", peopleId)
	headers := getHeaders()
	headers["Referer"] = fmt.Sprintf("https://www.zhihu.com/people/%s/answers", peopleId)

	if resp, err := g.Client().SetHeaderMap(headers).Get(answerGetUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.GetArray("data")

		// 获取用户名
		peopleName := jsonResp.GetString("data.0.author.name")
		if peopleName == "知乎用户" {
			activitiesUrl := fmt.Sprintf("https://www.zhihu.com/api/v4/members/%s/activities?limit=1", peopleId)
			if activitiesResp, err := g.Client().SetHeaderMap(headers).Get(activitiesUrl); err == nil {
				activitiesJsonResp := gjson.New(activitiesResp.ReadAllString())
				peopleName = activitiesJsonResp.GetString("data.0.actor.name")
			}
		}

		rssData := make(map[string]interface{})
		rssData["title"] = fmt.Sprintf("%s的知乎回答", peopleName)
		rssData["link"] = fmt.Sprintf("https://www.zhihu.com/people/%s/answers", peopleId)

		items := make([]map[string]string, 0)
		for index := range respDataList {
			title := jsonResp.GetString(fmt.Sprintf("data.%d.question.title", index))
			questionId := jsonResp.GetString(fmt.Sprintf("data.%d.question.id", index))
			questionAnswerId := jsonResp.GetString(fmt.Sprintf("data.%d.id", index))
			link := fmt.Sprintf("https://www.zhihu.com/question/%s/answer/%s", questionId, questionAnswerId)

			// 获取回答内容
			var description string
			answerDetailUrl := fmt.Sprintf("https://api.zhihu.com/appview/api/v4/answers/%s?include=content&is_appview=true", questionAnswerId)
			if answerResp, err := g.Client().SetHeaderMap(headers).Get(answerDetailUrl); err == nil {
				answerJsonResp := gjson.New(answerResp.ReadAllString())
				description = answerJsonResp.GetString("content")
				isError := answerJsonResp.GetString("error")
				if isError != "" {
					description = fmt.Sprintf("<a href='%s' target='_blank'>%s</a>", link, title)
				}
			} else {
				description = fmt.Sprintf("<a href='%s' target='_blank'>%s</a>", link, title)
			}

			itemMap := make(map[string]string)
			itemMap["title"] = title
			itemMap["description"] = description
			itemMap["author"] = peopleName
			timeStamp := jsonResp.GetInt64(fmt.Sprintf("data.%d.created_time", index))
			itemMap["pubDate"] = time.Unix(timeStamp, 0).String()
			itemMap["link"] = link
			items = append(items, itemMap)

		}

		rssData["items"] = items
		rssStr := lib.GenerateRSS(rssData)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
