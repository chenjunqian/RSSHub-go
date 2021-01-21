package dongqiudi

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
	"strings"
)

func (ctl *Controller) GetTopNews(req *ghttp.Request) {

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getNewsLinks()[linkType]

	cacheKey := "DONGQIUDI_TOP_NEWS_" + linkConfig.ChannelId
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := fmt.Sprintf("https://api.dongqiudi.com/app/tabs/iphone/%s.json?mark=gif&version=576", linkConfig.ChannelId)
	rssData := dao.RSSFeed{
		Title:       "懂球帝 - " + linkConfig.Title,
		Link:        apiUrl,
		Description: "懂球帝|专业权威的足球网站",
		ImageUrl:    "https://static1.dongqiudi.com/web-new/web/images/fav.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		respJson := gjson.New(resp.ReadAllString())
		articleJsonList := respJson.GetJsons("articles")
		rssItems := make([]dao.RSSItem, 0)
		for _, articleJson := range articleJsonList {
			title := articleJson.GetString("title")
			link := articleJson.GetString("url")
			imageLink := articleJson.GetString("thumb")
			time := articleJson.GetString("published_at")
			author := articleJson.GetString("author_classify")
			rssItem := dao.RSSItem{
				Title:       title,
				Link:        link,
				Description: fmt.Sprintf("<img src='%s'><br>", imageLink),
				Author:      author,
				Created:     time,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*6)
	_ = req.Response.WriteXmlExit(rssStr)
}
