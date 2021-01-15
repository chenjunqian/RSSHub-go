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

func (ctl *Controller) GetZhihuHostList(req *ghttp.Request) {
	redisKey := "ZHIHU_HOT_LIST"
	if value, err := g.Redis().DoVar("GET", redisKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	hotListUrl := "https://www.zhihu.com/api/v3/explore/guest/feeds?limit=40"
	headers := getHeaders()
	if resp, err := g.Client().SetHeaderMap(headers).Get(hotListUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.GetArray("data")

		rssData := dao.RSSFeed{}
		rssData.Title = "知乎热榜"
		rssData.Link = "https://www.zhihu.com/billboard"
		rssData.Description = "知乎热榜"
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"
		items := make([]dao.RSSItem, 0)
		for index := range respDataList {
			itemType := jsonResp.GetString(fmt.Sprintf("data.%d.target.type", index))
			feedItem := dao.RSSItem{}
			switch itemType {
			case "answer":
				questionId := jsonResp.GetString(fmt.Sprintf("data.%d.target.question.id", index))
				answerId := jsonResp.GetString(fmt.Sprintf("data.%d.target.id", index))
				content := jsonResp.GetString(fmt.Sprintf("data.%d.target.content", index))
				createdTime := jsonResp.GetInt64(fmt.Sprintf("data.%d.target.updated_time", index))

				feedItem.Title = jsonResp.GetString(fmt.Sprintf("data.%d.target.question.title", index))
				feedItem.Link = fmt.Sprintf("https://www.zhihu.com/%s/answer/%s", questionId, answerId)
				feedItem.Author = jsonResp.GetString(fmt.Sprintf("data.%d.target.author.name", index))
				feedItem.Description = fmt.Sprintf("%s的回答<br/><br/>%s", feedItem.Author, content)
				feedItem.Created = time.Unix(createdTime, 0).String()
			case "article":
				feedItem.Title = jsonResp.GetString(fmt.Sprintf("data.%d.target.title", index))
				content := jsonResp.GetString(fmt.Sprintf("data.%d.target.content", index))
				createdTime := jsonResp.GetInt64(fmt.Sprintf("data.%d.target.updated_time", index))
				articleId := jsonResp.GetString(fmt.Sprintf("data.%d.target.id", index))

				feedItem.Title = jsonResp.GetString(fmt.Sprintf("data.%d.target.title", index))
				feedItem.Link = fmt.Sprintf("https://zhuanlan.zhihu.com/p/%s", articleId)
				feedItem.Author = jsonResp.GetString(fmt.Sprintf("data.%d.target.author.name", index))
				feedItem.Description = fmt.Sprintf("%s的回答<br/><br/>%s", feedItem.Author, content)
				feedItem.Created = time.Unix(createdTime, 0).String()
			}

			items = append(items, feedItem)
		}

		rssData.Items = items
		rssStr := lib.GenerateRSS(rssData)
		g.Redis().DoVar("SET", redisKey, rssStr)
		g.Redis().DoVar("EXPIRE", redisKey, 60*60*1)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
