package pingwest

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
)

func (ctl *Controller) GetState(req *ghttp.Request) {

	cacheKey := "PINGWEST_STATE"
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.pingwest.com/api/state/list"
	rssData := dao.RSSFeed{
		Title:       "品玩 - 时事要闻",
		Link:        apiUrl,
		Tag:         []string{"时事", "科技"},
		Description: "品玩是具有全球化视野的科技内容平台和创新连接器，致力于服务全球科技创新者。",
		ImageUrl:    "https://cdn.pingwest.com/static/pingwest-logo-cn.jpg",
	}

	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := stateParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}

func stateParser(respString string) (items []dao.RSSItem) {
	respJson := gjson.New(respString)
	dataListHtml := respJson.GetString("data.list")
	dataSoup := soup.HTMLParse(dataListHtml)
	articleList := dataSoup.FindAll("section", "class", "item")
	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		time = gtime.Now().Format("Y-m-d") + " " + article.Find("section", "class", "time").Find("span").Text()
		title = article.Find("p", "class", "title").Find("a").Text()
		link = article.Find("p", "class", "title").Find("a").Attrs()["href"]
		if contentDoc := article.Find("p", "class", "description"); contentDoc.Error == nil {
			content = contentDoc.Find("a").Text()
		}

		if imageDoc := article.Find("section", "class", "news-img"); imageDoc.Error == nil {
			imageLink = imageDoc.Find("img").Attrs()["src"]
		}

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: feed.GenerateDescription(imageLink, content),
			Created:     time,
		}
		items = append(items, rssItem)
	}
	return
}
