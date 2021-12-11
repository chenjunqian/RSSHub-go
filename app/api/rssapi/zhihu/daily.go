package zhihu

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"regexp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
)

func (ctl *Controller) GetDaily(req *ghttp.Request) {

	if value, err := g.Redis().DoVar("GET", "ZHIHU_DAILY"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	dailyUrl := "https://news-at.zhihu.com/api/4/news/latest"
	headers := getHeaders()
	headers["Referer"] = dailyUrl
	if resp, err := component.GetHttpClient().SetHeaderMap(headers).Get(dailyUrl); err == nil {
		defer func(resp *ghttp.ClientResponse) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(err)
			}
		}(resp)
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.GetArray("stories")

		rssData := dao.RSSFeed{}
		rssData.Title = "知乎日报"
		rssData.Link = dailyUrl
		rssData.Description = "每天3次，每次7分钟"
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"

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
			key := fmt.Sprintf("ZHIHU_DAILY_%s", storyId)
			value, _ := g.Redis().DoVar("GET", key)
			if value.String() != "" {
				// 如果缓存里有就使用缓存内容
				feedItem.Description = value.String()
			} else {
				storyUrl := fmt.Sprintf("https://news-at.zhihu.com/api/4/news/%s", storyId)
				storyItemResp, err := component.GetHttpClient().SetHeaderMap(headers).Get(storyUrl)
				if err == nil {
					storyItemJsonResp := gjson.New(storyItemResp.ReadAllString())
					feedItem.Description = storyItemJsonResp.GetString("body")
					reg := regexp.MustCompile(`<div class="meta">([\s\S]*?)<\/div>`)
					feedItem.Description = reg.ReplaceAllString(feedItem.Description, `<strong>${1}</strong>`)
					reg = regexp.MustCompile(`<\/?h2.*?>`)
					feedItem.Description = reg.ReplaceAllString(feedItem.Description, "")
					_ = storyItemResp.Close()
				}
			}

			items = append(items, feedItem)
		}

		rssData.Items = items
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		g.Redis().DoVar("SET", "ZHIHU_DAILY", rssStr)
		g.Redis().DoVar("EXPIRE", "ZHIHU_DAILY", 60*60*1)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
