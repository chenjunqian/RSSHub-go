package chaping

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"
)

func (ctl *controller) GetNews(req *ghttp.Request) {

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getNewsLinks()[linkType]

	cacheKey := "CHAPING_NEWS_" + linkConfig.Caty
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	targetUrl := "https://chaping.cn/news?cate=" + linkConfig.Caty
	currentUrl := "https://chaping.cn/api/official/information/news?page=1&limit=16&cate=" + linkConfig.Caty
	rssData := dao.RSSFeed{
		Title:       "差评-" + linkConfig.Title,
		Link:        targetUrl,
		Description: "新媒体「差评」站在科技领域，以专业的角度和诙谐的语言，为订阅者呈现高质量的图文和视频内容",
		ImageUrl:    "https://chaping.cn/public/favicon.ico",
	}
	if resp, err := component.GetHttpClient().SetHeaderMap(getHeaders()).Get(currentUrl); err == nil {
		respJson := gjson.New(resp.ReadAllString())
		dataJsonList := respJson.GetJsons("data")
		rssItems := make([]dao.RSSItem, 0)
		for _, dataJson := range dataJsonList {
			title := dataJson.GetString("title")
			summary := dataJson.GetString("summary")
			imageLink := dataJson.GetString("images.url")
			time := dataJson.GetString("time_publish_timestamp")
			author := dataJson.GetString("author")
			link := dataJson.GetString("origin_url")
			fullContent := dataJson.GetString("content")
			rssItem := dao.RSSItem{
				Title:       title,
				Link:        link,
				Description: feed.GenerateDescription(imageLink, summary+"<br>"+fullContent),
				Author:      author,
				Created:     time,
				Thumbnail:   imageLink,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*3)
	_ = req.Response.WriteXmlExit(rssStr)
}
