package zhihu

import (
	"context"
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetTopic(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	topicId := req.Get("id").String()
	redisKey := "ZHIHU_TOPIC"
	if value, err := component.GetRedis().Do(ctx,"GET", redisKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	topicGetUrl := "https://www.zhihu.com/api/v4/topics/" + topicId + "/feeds/timeline_activity?include=data%5B%3F%28target.type%3Dtopic_sticky_module%29%5D.target.data%5B%3F%28target.type%3Danswer%29%5D.target.content%2Crelationship.is_authorized%2Cis_author%2Cvoting%2Cis_thanked%2Cis_nothelp%3Bdata%5B%3F%28target.type%3Dtopic_sticky_module%29%5D.target.data%5B%3F%28target.type%3Danswer%29%5D.target.is_normal%2Ccomment_count%2Cvoteup_count%2Ccontent%2Crelevant_info%2Cexcerpt.author.badge%5B%3F%28type%3Dbest_answerer%29%5D.topics%3Bdata%5B%3F%28target.type%3Dtopic_sticky_module%29%5D.target.data%5B%3F%28target.type%3Darticle%29%5D.target.content%2Cvoteup_count%2Ccomment_count%2Cvoting%2Cauthor.badge%5B%3F%28type%3Dbest_answerer%29%5D.topics%3Bdata%5B%3F%28target.type%3Dtopic_sticky_module%29%5D.target.data%5B%3F%28target.type%3Dpeople%29%5D.target.answer_count%2Carticles_count%2Cgender%2Cfollower_count%2Cis_followed%2Cis_following%2Cbadge%5B%3F%28type%3Dbest_answerer%29%5D.topics%3Bdata%5B%3F%28target.type%3Danswer%29%5D.target.annotation_detail%2Ccontent%2Crelationship.is_authorized%2Cis_author%2Cvoting%2Cis_thanked%2Cis_nothelp%3Bdata%5B%3F%28target.type%3Danswer%29%5D.target.author.badge%5B%3F%28type%3Dbest_answerer%29%5D.topics%3Bdata%5B%3F%28target.type%3Darticle%29%5D.target.annotation_detail%2Ccontent%2Cauthor.badge%5B%3F%28type%3Dbest_answerer%29%5D.topics%3Bdata%5B%3F%28target.type%3Dquestion%29%5D.target.annotation_detail%2Ccomment_count&limit=20"
	link := fmt.Sprintf("https://www.zhihu.com/topic/%s/newest", topicId)
	headers := getHeaders()
	headers["Authorization"] = "oauth c3cef7c66a1843f8b3a9e6a1e3160e20"
	headers["Referer"] = link
	if resp, err := component.GetHttpClient().SetHeaderMap(headers).Get(ctx, topicGetUrl); err == nil {
		defer func(resp *gclient.Response) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}(resp)
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.GetJsons("data")

		rssData := dao.RSSFeed{}
		rssData.Title = fmt.Sprintf("知乎话题-%s", topicId)
		rssData.Link = link
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"
		items := make([]dao.RSSItem, 0)
		for _, dataJson := range respDataList {
			var title string
			var description string
			var link string
			var pubDate string
			var author string

			targetJson := dataJson.GetJson("target")
			dataType := targetJson.Get("type").String()
			switch dataType {
			case "answer":
				title = fmt.Sprintf("%s-%s的回答：%s", targetJson.Get("question.title"), targetJson.Get("author.name"), targetJson.Get("excerpt"))
				description = fmt.Sprintf("<strong>%s</strong><br>%s的回答<br/>%s", targetJson.Get("question.title"), targetJson.Get("author.name"), targetJson.Get("content"))
				link = fmt.Sprintf("https://www.zhihu.com/question/%s/answer/%s", targetJson.Get("question.id"), targetJson.Get("id"))
				pubDate = time.Unix(targetJson.Get("updated_time").Int64(), 0).String()
				author = targetJson.Get("author.name").String()
			case "question":
				title = targetJson.Get("title").String()
				description = targetJson.Get("title").String()
				link = fmt.Sprintf("https://www.zhihu.com/question/%s", targetJson.Get("title"))
				pubDate = time.Unix(targetJson.Get("created").Int64(), 0).String()
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
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		req.Response.WriteXmlExit(rssStr)
	}
}
