package pingwest

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
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

	if resp := component.GetContent(apiUrl); resp != ""{

		rssItems := stateParser(resp)
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
		content = parseStateDetail(link)

		if imageDoc := article.Find("section", "class", "news-img"); imageDoc.Error == nil {
			imageLink = imageDoc.Find("img").Attrs()["src"]
		}

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: feed.GenerateDescription(imageLink, content),
			Created:     time,
			Thumbnail:   imageLink,
		}
		items = append(items, rssItem)
	}
	return
}

func parseStateDetail(detailLink string) (detailData string) {
	var (
		resp *ghttp.ClientResponse
		err  error
	)
	if strings.HasPrefix(detailLink, "//") {
		detailLink = "https:" + detailLink
	}
	if resp, err = component.GetHttpClient().SetHeaderMap(getHeaders()).Get(detailLink); err == nil {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)
		defer func(resp *ghttp.ClientResponse) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(err)
			}
		}(resp)
		respString = resp.ReadAllString()
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("section", "class", "wire-detail-box")
		if articleElem.Pointer == nil {
			articleElem = docs.Find("section", "class", "main")
		}
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request pingwest article detail failed, link  %s \n", detailLink)
	}

	return
}
