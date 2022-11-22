package huxiu

import (
	"context"
	"fmt"
	"regexp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetEvent(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	cacheKey := "HUXIU_EVENT"
	if value, err := component.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.huxiu.com/event/"
	rssData := dao.RSSFeed{
		Title:       "虎嗅 - 号外",
		Link:        apiUrl,
		Tag:         []string{"商业", "科技", "互联网"},
		Description: "聚合优质的创新信息与人群，捕获精选|深度|犀利的商业科技资讯。在虎嗅，不错过互联网的每个重要时刻。",
		ImageUrl:    "https://www.huxiu.com/favicon.ico",
	}
	if resp := component.GetContent(ctx,apiUrl); resp != "" {
		rssItems := make([]dao.RSSItem, 0)
		reg := regexp.MustCompile(`window.__INITIAL_STATE__=(.*?);\(function\(\)`)
		contentStrs := reg.FindStringSubmatch(resp)
		if len(contentStrs) <= 1 {
			return
		}
		contentStr := contentStrs[1]
		dataJsonList := gjson.New(contentStr).GetJsons("event.eventList.datalist")
		for _, dataJson := range dataJsonList {
			var imageLink string
			var title string
			var link string
			var author string
			var content string
			var time string

			title = dataJson.Get("name").String()
			time = dataJson.Get("publish_time").String()
			link = fmt.Sprintf("https://www.huxiu.com/event/%s.html", dataJson.Get("event_id"))
			imageLink, _ = gurl.Decode(dataJson.Get("cover_path").String())
			content = dataJson.Get("introduce").String()

			rssItem := dao.RSSItem{
				Title:     title,
				Link:      link,
				Author:    author,
				Content:   feed.GenerateContent(content),
				Created:   time,
				Thumbnail: imageLink,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	component.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
