package pingwest

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

func (ctl *Controller) GetState(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	cacheKey := "PINGWEST_STATE"
	if value, err := service.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
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

	if resp := service.GetContent(ctx,apiUrl); resp != "" {

		rssItems := stateParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	service.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	service.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}

func stateParser(ctx context.Context, respString string) (items []dao.RSSItem) {
	respJson := gjson.New(respString)
	dataListHtml := respJson.Get("data.list").String()
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
		content = parseStateDetail(ctx, link)

		if imageDoc := article.Find("section", "class", "news-img"); imageDoc.Error == nil {
			imageLink = imageDoc.Find("img").Attrs()["src"]
		}

		rssItem := dao.RSSItem{
			Title:     title,
			Link:      link,
			Author:    author,
			Content:   feed.GenerateContent(content),
			Created:   time,
			Thumbnail: imageLink,
		}
		items = append(items, rssItem)
	}
	return
}

func parseStateDetail(ctx context.Context, detailLink string) (detailData string) {
	var (
		resp *gclient.Response
		err  error
	)
	if strings.HasPrefix(detailLink, "//") {
		detailLink = "https:" + detailLink
	}
	if resp, err = service.GetHttpClient().SetHeaderMap(getHeaders()).Get(ctx, detailLink); err == nil {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)
		defer func(resp *gclient.Response) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(ctx, err)
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
		g.Log().Errorf(ctx,"Request pingwest article detail failed, link  %s \n", detailLink)
	}

	return
}
