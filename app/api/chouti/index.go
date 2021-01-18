package chouti

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
	"strings"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getNewsLinks()[linkType]

	cacheKey := "CHOUTI_" + linkConfig.LinkType
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	var apiUrl string
	if linkConfig.LinkType == "hot" {
		apiUrl = "https://m.chouti.com/api/m/link/hot?afterTime=0"
	} else {
		apiUrl = fmt.Sprintf("https://m.chouti.com/api/m/zone/%s?afterScore=0", linkConfig.LinkType)
	}
	rssData := dao.RSSFeed{
		Title:       "抽屉新热榜" + linkConfig.Title,
		Link:        "https://m.chouti.com",
		Description: "抽屉新热榜，汇聚每日搞笑段子、热门图片、有趣新闻。它将微博、门户、社区、bbs、社交网站等海量内容聚合在一起，通过用户推荐生成最热榜单。看抽屉新热榜，每日热门、有趣资讯尽收眼底",
		ImageUrl:    "https://m.chouti.com/static/image/favicon.png",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		respJson := gjson.New(resp.ReadAllString())
		dataJsonList := respJson.GetJsons("data")
		rssItems := make([]dao.RSSItem, 0)
		for _, dataJson := range dataJsonList {
			title := dataJson.GetString("title")
			imageLink := dataJson.GetString("img_url")
			time := dataJson.GetString("createTime")
			author := dataJson.GetString("submitted_user.nick")
			link := dataJson.GetString("originalUrl")
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
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*3)
	_ = req.Response.WriteXmlExit(rssStr)
}
