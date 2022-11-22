package zhihu

import (
	"context"
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetZhihuHostList(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	redisKey := "ZHIHU_HOT_LIST"
	if value, err := component.GetRedis().Do(ctx,"GET", redisKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	hotListUrl := "https://www.zhihu.com/api/v3/explore/guest/feeds?limit=40"
	headers := getHeaders()
	if resp, err := component.GetHttpClient().SetHeaderMap(headers).Get(ctx, hotListUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.Get("data").Strings()

		rssData := dao.RSSFeed{}
		rssData.Title = "知乎热榜"
		rssData.Link = "https://www.zhihu.com/billboard"
		rssData.Description = "知乎热榜"
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"
		items := make([]dao.RSSItem, 0)
		for index := range respDataList {
			itemType := jsonResp.Get(fmt.Sprintf("data.%d.target.type", index)).String()
			feedItem := dao.RSSItem{}
			switch itemType {
			case "answer":
				questionId := jsonResp.Get(fmt.Sprintf("data.%d.target.question.id", index))
				answerId := jsonResp.Get(fmt.Sprintf("data.%d.target.id", index))
				content := jsonResp.Get(fmt.Sprintf("data.%d.target.content", index))
				createdTime := jsonResp.Get(fmt.Sprintf("data.%d.target.updated_time", index)).Int64()

				feedItem.Title = jsonResp.Get(fmt.Sprintf("data.%d.target.question.title", index)).String()
				feedItem.Link = fmt.Sprintf("https://www.zhihu.com/%s/answer/%s", questionId, answerId)
				feedItem.Author = jsonResp.Get(fmt.Sprintf("data.%d.target.author.name", index)).String()
				feedItem.Description = fmt.Sprintf("%s的回答<br/><br/>%s", feedItem.Author, content)
				feedItem.Created = time.Unix(createdTime, 0).String()
			case "article":
				feedItem.Title = jsonResp.Get(fmt.Sprintf("data.%d.target.title", index)).String()
				content := jsonResp.Get(fmt.Sprintf("data.%d.target.content", index))
				createdTime := jsonResp.Get(fmt.Sprintf("data.%d.target.updated_time", index)).Int64()
				articleId := jsonResp.Get(fmt.Sprintf("data.%d.target.id", index))

				feedItem.Title = jsonResp.Get(fmt.Sprintf("data.%d.target.title", index)).String()
				feedItem.Link = fmt.Sprintf("https://zhuanlan.zhihu.com/p/%s", articleId)
				feedItem.Author = jsonResp.Get(fmt.Sprintf("data.%d.target.author.name", index)).String()
				feedItem.Description = fmt.Sprintf("%s的回答<br/><br/>%s", feedItem.Author, content)
				feedItem.Created = time.Unix(createdTime, 0).String()
			}

			items = append(items, feedItem)
		}

		rssData.Items = items
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		component.GetRedis().Do(ctx,"SET", redisKey, rssStr)
		component.GetRedis().Do(ctx,"EXPIRE", redisKey, 60*60*1)
		req.Response.WriteXmlExit(rssStr)
	}
}
