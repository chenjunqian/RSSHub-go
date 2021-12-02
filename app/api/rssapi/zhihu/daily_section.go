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

func (ctl *Controller) GetZhihuDailySection(req *ghttp.Request) {

	sectionId := req.Get("id")
	redisKey := fmt.Sprintf("ZHIHU_DAILY_SECTION_%s", sectionId)
	if value, err := g.Redis().DoVar("GET", redisKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	dailyUrl := fmt.Sprintf("https://news-at.zhihu.com/api/7/section/%s", sectionId)
	headers := getHeaders()
	headers["Referer"] = dailyUrl

	if resp, err := component.GetHttpClient().SetHeaderMap(headers).Get(dailyUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.GetArray("stories")

		rssData := dao.RSSFeed{}
		rssData.Title = "知乎日报"
		rssData.Link = "https://daily.zhihu.com"
		rssData.Description = "每天3次，每次7分钟"
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"
		items := make([]dao.RSSItem, 0)
		for index := range respDataList {
			feedItem := dao.RSSItem{
				Title: jsonResp.GetString(fmt.Sprintf("stories.%d.title", index)),
				Link:  jsonResp.GetString(fmt.Sprintf("stories.%d.url", index)),
			}
			storyId := jsonResp.GetString(fmt.Sprintf("stories.%d.id", index))
			storyUrl := fmt.Sprintf("https://news-at.zhihu.com/api/7/news/%s", storyId)
			if itemResp, err := component.GetHttpClient().SetHeaderMap(headers).Get(storyUrl); err == nil {
				jsonItemResp := gjson.New(itemResp.ReadAllString())
				content := jsonItemResp.GetString("body")
				reg := regexp.MustCompile(`<div class="meta">([\s\S]*?)<\/div>`)
				content = reg.ReplaceAllString(content, `<strong>${1}</strong>`)
				reg = regexp.MustCompile(`<\/?h2.*?>`)
				content = reg.ReplaceAllString(content, "")
				feedItem.Description = content
			}
			items = append(items, feedItem)
		}

		rssData.Items = items
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		g.Redis().DoVar("SET", redisKey, rssStr)
		g.Redis().DoVar("EXPIRE", redisKey, 60*60*6)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
