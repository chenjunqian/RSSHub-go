package zhihu

import (
	"fmt"
	"rsshub/app/dao"
	"rsshub/lib"
	"time"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetTopic(req *ghttp.Request) {
	topicId := req.GetString("id")
	redisKey := "ZHIHU_TOPIC"
	if value, err := g.Redis().DoVar("GET", redisKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value)
		}
	}

	topicGetUrl := "https://www.zhihu.com/api/v4/topics/" + topicId + "/feeds/timeline_activity?include=data%5B%3F%28target.type%3Dtopic_sticky_module%29%5D.target.data%5B%3F%28target.type%3Danswer%29%5D.target.content%2Crelationship.is_authorized%2Cis_author%2Cvoting%2Cis_thanked%2Cis_nothelp%3Bdata%5B%3F%28target.type%3Dtopic_sticky_module%29%5D.target.data%5B%3F%28target.type%3Danswer%29%5D.target.is_normal%2Ccomment_count%2Cvoteup_count%2Ccontent%2Crelevant_info%2Cexcerpt.author.badge%5B%3F%28type%3Dbest_answerer%29%5D.topics%3Bdata%5B%3F%28target.type%3Dtopic_sticky_module%29%5D.target.data%5B%3F%28target.type%3Darticle%29%5D.target.content%2Cvoteup_count%2Ccomment_count%2Cvoting%2Cauthor.badge%5B%3F%28type%3Dbest_answerer%29%5D.topics%3Bdata%5B%3F%28target.type%3Dtopic_sticky_module%29%5D.target.data%5B%3F%28target.type%3Dpeople%29%5D.target.answer_count%2Carticles_count%2Cgender%2Cfollower_count%2Cis_followed%2Cis_following%2Cbadge%5B%3F%28type%3Dbest_answerer%29%5D.topics%3Bdata%5B%3F%28target.type%3Danswer%29%5D.target.annotation_detail%2Ccontent%2Crelationship.is_authorized%2Cis_author%2Cvoting%2Cis_thanked%2Cis_nothelp%3Bdata%5B%3F%28target.type%3Danswer%29%5D.target.author.badge%5B%3F%28type%3Dbest_answerer%29%5D.topics%3Bdata%5B%3F%28target.type%3Darticle%29%5D.target.annotation_detail%2Ccontent%2Cauthor.badge%5B%3F%28type%3Dbest_answerer%29%5D.topics%3Bdata%5B%3F%28target.type%3Dquestion%29%5D.target.annotation_detail%2Ccomment_count&limit=20"
	link := fmt.Sprintf("https://www.zhihu.com/topic/%s/newest", topicId)
	headers := getHeaders()
	headers["Authorization"] = "oauth c3cef7c66a1843f8b3a9e6a1e3160e20"
	headers["Referer"] = link
	fmt.Println("topicGetUrl : ", topicGetUrl)
	if resp, err := g.Client().SetHeaderMap(headers).Get(topicGetUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.GetJsons("data")

		rssData := dao.RSSFeed{}
		rssData.Title = fmt.Sprintf("知乎话题-%s", topicId)
		rssData.Link = link
		items := make([]dao.RSSItem, 0)
		for _, dataJson := range respDataList {
			var title string
			var description string
			var link string
			var pubDate string
			var author string

			targetJson := dataJson.GetJson("target")
			dataType := targetJson.GetString("type")
			switch dataType {
			case "answer":
				title = fmt.Sprintf("%s-%s的回答：%s", targetJson.GetString("question.title"), targetJson.GetString("author.name"), targetJson.GetString("excerpt"))
				description = fmt.Sprintf("<strong>%s</strong><br>%s的回答<br/>%s", targetJson.GetString("question.title"), targetJson.GetString("author.name"), targetJson.GetString("content"))
				link = fmt.Sprintf("https://www.zhihu.com/question/%s/answer/%s", targetJson.GetString("question.id"), targetJson.GetString("id"))
				pubDate = time.Unix(targetJson.GetInt64("updated_time"), 0).String()
				author = targetJson.GetString("author.name")
			case "question":
				title = targetJson.GetString("title")
				description = targetJson.GetString("title")
				link = fmt.Sprintf("https://www.zhihu.com/question/%s", targetJson.GetString("title"))
				pubDate = time.Unix(targetJson.GetInt64("created"), 0).String()
			}

			rssItem := dao.RSSItem{
				Title:       title,
				Description: description,
				Author:      author,
				Link:        link,
				Created:     pubDate,
			}
			items = append(items, rssItem)
		}

		rssData.Items = items
		rssStr := lib.GenerateRSS(rssData)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
