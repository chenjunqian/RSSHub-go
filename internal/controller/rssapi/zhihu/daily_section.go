package zhihu

import (
	"context"
	"fmt"
	"regexp"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetZhihuDailySection(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	sectionId := req.Get("id")
	redisKey := fmt.Sprintf("ZHIHU_DAILY_SECTION_%s", sectionId)
	if value, err := service.GetRedis().Do(ctx,"GET", redisKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	dailyUrl := fmt.Sprintf("https://news-at.zhihu.com/api/7/section/%s", sectionId)
	headers := getHeaders()
	headers["Referer"] = dailyUrl

	if resp, err := service.GetHttpClient().SetHeaderMap(headers).Get(ctx, dailyUrl); err == nil {
		defer func(resp *gclient.Response) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}(resp)
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.Get("stories").Strings()

		rssData := dao.RSSFeed{}
		rssData.Title = "知乎日报"
		rssData.Link = "https://daily.zhihu.com"
		rssData.Description = "每天3次，每次7分钟"
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"
		items := make([]dao.RSSItem, 0)
		for index := range respDataList {
			feedItem := dao.RSSItem{
				Title: jsonResp.Get(fmt.Sprintf("stories.%d.title", index)).String(),
				Link:  jsonResp.Get(fmt.Sprintf("stories.%d.url", index)).String(),
			}
			storyId := jsonResp.Get(fmt.Sprintf("stories.%d.id", index))
			storyUrl := fmt.Sprintf("https://news-at.zhihu.com/api/7/news/%s", storyId)
			if itemResp, err := service.GetHttpClient().SetHeaderMap(headers).Get(ctx, storyUrl); err == nil {
				jsonItemResp := gjson.New(itemResp.ReadAllString())
				content := jsonItemResp.Get("body").String()
				reg := regexp.MustCompile(`<div class="meta">([\s\S]*?)<\/div>`)
				content = reg.ReplaceAllString(content, `<strong>${1}</strong>`) 
				reg = regexp.MustCompile(`<\/?h2.*?>`)
				content = reg.ReplaceAllString(content, "")
				feedItem.Description = content
				_ = itemResp.Close()
			}
			items = append(items, feedItem)
		}

		rssData.Items = items
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		service.GetRedis().Do(ctx,"SET", redisKey, rssStr)
		service.GetRedis().Do(ctx,"EXPIRE", redisKey, 60*60*6)
		req.Response.WriteXmlExit(rssStr)
	}
}
