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

	if resp, err := g.Client().SetHeaderMap(headers).Get(dailyUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.GetArray("stories")

		rssData := dao.RSSFeed{}
		rssData.Title = "知乎日报"
		rssData.Link = "https://daily.zhihu.com"
		rssData.Description = "每天3次，每次7分钟"
		items := make([]dao.RSSItem, 0)
		for index := range respDataList {
			feedItem := dao.RSSItem{
				Title: jsonResp.GetString(fmt.Sprintf("stories.%d.title", index)),
				Link:  jsonResp.GetString(fmt.Sprintf("stories.%d.url", index)),
			}
			storyId := jsonResp.GetString(fmt.Sprintf("stories.%d.id", index))
			storyUrl := fmt.Sprintf("https://news-at.zhihu.com/api/7/news/%s", storyId)
			if itemResp, err := g.Client().SetHeaderMap(headers).Get(storyUrl); err == nil {
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
		rssStr := lib.GenerateRSS(rssData)
		g.Redis().DoVar("SET", redisKey, rssStr)
		g.Redis().DoVar("EXPIRE", redisKey, 60*60*6)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
