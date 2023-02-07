package huxiu

import (
	"context"
	"fmt"
	"regexp"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetTimeline(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	cacheKey := "HUXIU_TIMELINE"
	if value, err := service.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.huxiu.com/timeline/"
	rssData := dao.RSSFeed{
		Title:       "虎嗅 - 时间线",
		Link:        apiUrl,
		Tag:         []string{"商业", "科技", "互联网"},
		Description: "聚合优质的创新信息与人群，捕获精选|深度|犀利的商业科技资讯。在虎嗅，不错过互联网的每个重要时刻。",
		ImageUrl:    "https://www.huxiu.com/favicon.ico",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != "" {
		rssItems := make([]dao.RSSItem, 0)
		reg := regexp.MustCompile(`window.__INITIAL_STATE__=(.*?);\(function\(\)`)
		contentStrs := reg.FindStringSubmatch(resp)
		if len(contentStrs) <= 1 {
			return
		}
		contentStr := contentStrs[1]
		dataJsonList := gjson.New(contentStr).GetJsons("category.timelineList.datalist")
		for _, dataJson := range dataJsonList {
			var imageLink string
			var title string
			var link string
			var author string
			var content string
			var time string

			title = dataJson.Get("name").String()
			time = dataJson.Get("update_time").String()
			link = fmt.Sprintf("https://www.huxiu.com/timeline/%s.html", dataJson.Get("id"))
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
	service.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	service.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
