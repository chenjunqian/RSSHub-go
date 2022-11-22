package houxu

import (
	"context"
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	cacheKey := "HOUXU_INDEX"
	if value, err := component.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://houxu.app/api/1/bundle/index/"
	rssData := dao.RSSFeed{
		Title:       "后续 - 热点",
		Link:        apiUrl,
		Tag:         []string{"互联网"},
		Description: "后续 · 有记忆的新闻，持续追踪热点新闻",
		ImageUrl:    "https://assets-1256259474.cos.ap-shanghai.myqcloud.com/static/img/icon-180.jpg",
	}
	if resp := component.GetContent(ctx,apiUrl); resp != "" {

		respJson := gjson.New(resp)
		dataJsonArray := respJson.GetJsons("indexRecords.results")
		rssItems := make([]dao.RSSItem, 0)
		for _, dataJson := range dataJsonArray {
			var imageLink string
			var title string
			var link string
			var author string
			var content string
			var time string

			title = dataJson.Get("object.title").String()
			link = fmt.Sprintf("https://houxu.app/lives/%s", dataJson.Get("object.id"))
			author = dataJson.Get("creator.name").String()
			content = dataJson.Get("object.summary").String()
			time = dataJson.Get("object.news_update_at").String()
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
