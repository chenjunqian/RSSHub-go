package huxiu

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gurl"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"regexp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
)

func (ctl *Controller) GetCollection(req *ghttp.Request) {

	cacheKey := "HUXIU_COLLECTION"
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.huxiu.com/collection/"
	rssData := dao.RSSFeed{
		Title:       "虎嗅 - 文集",
		Link:        apiUrl,
		Tag:         []string{"文学"},
		Description: "聚合优质的创新信息与人群，捕获精选|深度|犀利的商业科技资讯。在虎嗅，不错过互联网的每个重要时刻。",
		ImageUrl:    "https://www.huxiu.com/favicon.ico",
	}
	if resp, err := component.GetHttpClient().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := make([]dao.RSSItem, 0)
		reg := regexp.MustCompile(`window.__INITIAL_STATE__=(.*?);\(function\(\)`)
		contentStrs := reg.FindStringSubmatch(resp.ReadAllString())
		if len(contentStrs) <= 1 {
			return
		}
		contentStr := contentStrs[1]
		dataJsonList := gjson.New(contentStr).GetJsons("category.collectionList.datalist")
		for _, dataJson := range dataJsonList {
			var imageLink string
			var title string
			var link string
			var author string
			var content string
			var time string

			title = dataJson.GetString("name")
			time = dataJson.GetString("update_time")
			link = fmt.Sprintf("https://www.huxiu.com/collection/%s.html", dataJson.GetString("id"))
			imageLink, _ = gurl.Decode(dataJson.GetString("head_img"))
			content = parseCommonDetail(link)

			rssItem := dao.RSSItem{
				Title:       title,
				Link:        link,
				Author:      author,
				Description: feed.GenerateDescription(imageLink, content),
				Created:     time,
				Thumbnail:   imageLink,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
