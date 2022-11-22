package zhihu

import (
	"context"
	"fmt"
	"regexp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetDaily(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	if value, err := component.GetRedis().Do(ctx,"GET", "ZHIHU_DAILY"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	dailyUrl := "https://news-at.zhihu.com/api/4/news/latest"
	headers := getHeaders()
	headers["Referer"] = dailyUrl
	if resp, err := component.GetHttpClient().SetHeaderMap(headers).Get(ctx, dailyUrl); err == nil {
		defer func(resp *gclient.Response) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(ctx ,err)
			}
		}(resp)
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.Get("stories").Strings()

		rssData := dao.RSSFeed{}
		rssData.Title = "知乎日报"
		rssData.Link = dailyUrl
		rssData.Description = "每天3次，每次7分钟"
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"

		items := make([]dao.RSSItem, 0)
		for index := range respDataList {

			feedItem := dao.RSSItem{
				Title: jsonResp.Get(fmt.Sprintf("stories.%d.title", index)).String(),
				Link:  jsonResp.Get(fmt.Sprintf("stories.%d.url", index)).String(),
			}

			storyType := jsonResp.Get(fmt.Sprintf("stories.%d.type", index)).String()
			if storyType == "1" {
				// 根据api的说明，过滤掉极个别站外链接
				continue
			}
			storyId := jsonResp.Get(fmt.Sprintf("stories.%d.id", index))
			key := fmt.Sprintf("ZHIHU_DAILY_%s", storyId)
			value, _ := component.GetRedis().Do(ctx,"GET", key)
			if value.String() != "" {
				// 如果缓存里有就使用缓存内容
				feedItem.Description = value.String()
			} else {
				storyUrl := fmt.Sprintf("https://news-at.zhihu.com/api/4/news/%s", storyId)
				storyItemResp, err := component.GetHttpClient().SetHeaderMap(headers).Get(ctx, storyUrl)
				if err == nil {
					storyItemJsonResp := gjson.New(storyItemResp.ReadAllString())
					feedItem.Description = storyItemJsonResp.Get("body").String()
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
		component.GetRedis().Do(ctx,"SET", "ZHIHU_DAILY", rssStr)
		component.GetRedis().Do(ctx,"EXPIRE", "ZHIHU_DAILY", 60*60*1)
		req.Response.WriteXmlExit(rssStr)
	}
}
