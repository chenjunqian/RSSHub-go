package zhihu

import (
	"fmt"
	"rsshub/lib"
	"strconv"
	"time"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetActivities(req *ghttp.Request) {
	id := req.Get("id")
	url := fmt.Sprintf("https://www.zhihu.com/api/v4/members/%s/activities?limit=7", id)
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(url); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList, ok := jsonResp.Get("data").([]interface{})
		if ok {
			rssData := make(map[string]interface{})
			mainTitle := jsonResp.GetString("data.0.actor.name")
			mainDescription := jsonResp.GetString("data.0.actor.headline")
			if mainDescription != "" {
				mainDescription = jsonResp.GetString("data.0.actor.description")
			}
			rssData["title"] = mainTitle
			rssData["description"] = mainDescription
			rssData["link"] = url

			items := make([]map[string]string, 0)
			for index := range respDataList {
				var title string
				var description string
				var url string
				var author string
				contentType := jsonResp.GetString(fmt.Sprintf("data.%d.target.type", index))
				switch contentType {

				case "answer":
					title = jsonResp.GetString(fmt.Sprintf("data.%d.target.question.title", index))
					author = jsonResp.GetString(fmt.Sprintf("data.%d.target.author.name", index))
					// TODO format the content with images
					description = jsonResp.GetString(fmt.Sprintf("data.%d.target.content", index))
					questionId := jsonResp.GetFloat64(fmt.Sprintf("data.%d.target.question.id", index))
					answerId := jsonResp.GetFloat64(fmt.Sprintf("data.%d.target.id", index))
					url = fmt.Sprint("https://www.zhihu.com/question/$s/answer/$s", questionId, answerId)
				case "article":
				case "pin":
				case "question":
				case "collection":
				case "column":
				case "topic":
				case "live":
				case "roundtable":

				}
				itemMap := make(map[string]string)
				itemMap["title"] = title
				itemMap["description"] = description
				itemMap["author"] = author
				timeStampStr := jsonResp.GetFloat64(fmt.Sprintf("data.%d.created_time", index))
				timeStamp, _ := strconv.ParseInt(fmt.Sprintf("%f000", timeStampStr), 10, 64)
				itemMap["pubDate"] = time.Unix(timeStamp, 0).String()
				itemMap["url"] = url
				items = append(items, itemMap)
			}

			rssData["items"] = items
			rssStr := lib.GenerateRSS(rssData)
			_ = req.Response.WriteXmlExit(rssStr)
		}

	} else {
		_ = req.Response.WriteXmlExit("")
	}

}
