package bilibili

import (
	"context"
	"fmt"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetUserAudio(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	userId := req.Get("id").String()
	audioUrl := "https://www.bilibili.com/audio/au"
	link := "https://www.bilibili.com/audio/am" + userId

	rssData := dao.RSSFeed{}
	rssData.Link = link
	rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

	apiMenuUrl := "https://www.bilibili.com/audio/music-service-c/web/menu/info?sid=" + userId
	header := getHeaders()
	header["Referer"] = "https://space.bilibili.com/" + userId
	if menuResp, err := service.GetHttpClient().SetHeaderMap(header).SetCookieMap(getCookieMap(ctx)).Get(ctx, apiMenuUrl); err == nil {
		menuJsonResp := gjson.New(menuResp.ReadAllString())
		dataJson := menuJsonResp.GetJson("data")
		rssData.Description = dataJson.Get("intro").String()
		rssData.Title = dataJson.Get("title").String()
	}

	apiUrl := fmt.Sprintf("https://www.bilibili.com/audio/music-service-c/web/song/of-menu?sid=%s&pn=1&ps=100", userId)
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		jsonResp := gjson.New(resp)
		dataJsons := jsonResp.GetJsons("data.data")

		items := make([]dao.RSSItem, 0)
		for _, dataJson := range dataJsons {
			rssItem := dao.RSSItem{}
			rssItem.Title = dataJson.Get("title").String()
			rssItem.Link = audioUrl + dataJson.Get("statistic.sid").String()
			rssItem.Author = dataJson.Get("author").String()
			intro := dataJson.Get("intro").String()
			cover := dataJson.Get("cover").String()
			rssItem.Content = feed.GenerateContent(intro)
			rssItem.Thumbnail = cover
			items = append(items, rssItem)
		}

		rssData.Items = items
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		req.Response.WriteXmlExit(rssStr)
	}
}
