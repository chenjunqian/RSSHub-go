package chouti

import (
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *controller) GetIndex(req *ghttp.Request) {

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
	var rssData = dao.RSSFeed{
		Title:       "抽屉新热榜" + linkConfig.Title,
		Link:        "https://m.chouti.com",
		Tag:         linkConfig.Tags,
		Description: "抽屉新热榜，汇聚每日搞笑段子、热门图片、有趣新闻。它将微博、门户、社区、bbs、社交网站等海量内容聚合在一起，通过用户推荐生成最热榜单。看抽屉新热榜，每日热门、有趣资讯尽收眼底",
		ImageUrl:    "https://m.chouti.com/static/image/favicon.png",
	}
	if resp := component.GetContent(apiUrl); resp != "" {
		respJson := gjson.New(resp)
		dataJsonList := respJson.GetJsons("data")
		rssItems := make([]dao.RSSItem, 0)
		for _, dataJson := range dataJsonList {
			var (
				title     string
				imageLink string
				time      string
				author    string
				link      string
				content   string
			)
			title = dataJson.GetString("title")
			imageLink = dataJson.GetString("img_url")
			time = dataJson.GetString("createTime")
			author = dataJson.GetString("submitted_user.nick")
			link = dataJson.GetString("originalUrl")
			content = parseIndexDetail(link)
			rssItem := dao.RSSItem{
				Title:       title,
				Link:        link,
				Description: feed.GenerateDescription(imageLink, content),
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

func parseIndexDetail(detailLink string) (detailData string) {
	var (
		resp string
	)
	if resp = component.GetContent(detailData); resp != "" {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)
		respString = resp
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("div", "id", "container")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request chouti index article detail failed, link  %s \n", detailLink)
	}

	return
}
