package houxu

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {

	cacheKey := "HOUXU_INDEX"
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
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
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		respJson := gjson.New(resp.ReadAllString())
		dataJsonArray := respJson.GetJsons("indexRecords.results")
		rssItems := make([]dao.RSSItem, 0)
		for _, dataJson := range dataJsonArray {
			var imageLink string
			var title string
			var link string
			var author string
			var content string
			var time string

			title = dataJson.GetString("object.title")
			link = fmt.Sprintf("https://houxu.app/lives/%s", dataJson.GetString("object.id"))
			author = dataJson.GetString("creator.name")
			content = dataJson.GetString("object.summary")
			time = dataJson.GetString("object.news_update_at")
			rssItem := dao.RSSItem{
				Title:       title,
				Link:        link,
				Author:      author,
				Description: lib.GenerateDescription(imageLink, content),
				Created:     time,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
