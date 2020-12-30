package zhihu

import (
	"fmt"
	"regexp"
	"rsshub/app/dao"
	"rsshub/lib"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetDaily(req *ghttp.Request) {

	dailyUrl := "https://news-at.zhihu.com/api/4/news/latest"
	headers := getHeaders()
	headers["Referer"] = dailyUrl
	if resp, err := g.Client().SetHeaderMap(headers).Get(dailyUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.GetArray("stories")

		rssData := dao.RSSFeed{}
		rssData.Title = "知乎日报"
		rssData.Link = dailyUrl
		rssData.Description = "每天3次，每次7分钟"

		items := make([]dao.RSSItem, 0)
		for index := range respDataList {

			feedItem := dao.RSSItem{
				Title: jsonResp.GetString(fmt.Sprintf("stories.%d.title", index)),
				Link:  jsonResp.GetString(fmt.Sprintf("stories.%d.url", index)),
			}

			storyType := jsonResp.GetString(fmt.Sprintf("stories.%d.type", index))
			if storyType == "1" {
				// 根据api的说明，过滤掉极个别站外链接
				continue
			}
			storyId := jsonResp.GetString(fmt.Sprintf("stories.%d.id", index))
			key := fmt.Sprintf("zhihu_daily_%s", storyId)
			value, _ := g.Redis().DoVar("GET", key)
			if value.String() != "" {
				// 如果缓存里有就使用缓存内容
				feedItem.Description = value.String()
			} else {
				storyUrl := fmt.Sprintf("https://news-at.zhihu.com/api/4/news/%s", storyId)
				storyItemResp, err := g.Client().SetHeaderMap(headers).Get(storyUrl)
				if err == nil {
					storyItemJsonResp := gjson.New(storyItemResp.ReadAllString())
					feedItem.Description = storyItemJsonResp.GetString("body")
					reg := regexp.MustCompile(`<div class="meta">([\s\S]*?)<\/div>`)
					feedItem.Description = reg.ReplaceAllString(feedItem.Description, `<strong>${1}</strong>`)
					reg = regexp.MustCompile(`<\/?h2.*?>`)
					feedItem.Description = reg.ReplaceAllString(feedItem.Description, "")
					g.Redis().DoVar("SET", key, feedItem.Description)
				}
			}

			items = append(items, feedItem)
		}

		rssData.Items = items
		rssStr := lib.GenerateRSS(rssData)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
