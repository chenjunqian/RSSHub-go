package bilibili

import (
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetUserAudio(req *ghttp.Request) {
	userId := req.GetString("id")
	audioUrl := "https://www.bilibili.com/audio/au"
	link := "https://www.bilibili.com/audio/am" + userId

	rssData := dao.RSSFeed{}
	rssData.Link = link
	rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

	apiMenuUrl := "https://www.bilibili.com/audio/music-service-c/web/menu/info?sid=" + userId
	header := getHeaders()
	header["Referer"] = "https://space.bilibili.com/" + userId
	if menuResp, err := component.GetHttpClient().SetHeaderMap(header).SetCookieMap(getCookieMap()).Get(apiMenuUrl); err == nil {
		menuJsonResp := gjson.New(menuResp.ReadAllString())
		dataJson := menuJsonResp.GetJson("data")
		rssData.Description = dataJson.GetString("intro")
		rssData.Title = dataJson.GetString("title")
	}

	apiUrl := fmt.Sprintf("https://www.bilibili.com/audio/music-service-c/web/song/of-menu?sid=%s&pn=1&ps=100", userId)
	if resp := component.GetContent(apiUrl); resp != "" {
		jsonResp := gjson.New(resp)
		dataJsons := jsonResp.GetJsons("data.data")

		items := make([]dao.RSSItem, 0)
		for _, dataJson := range dataJsons {
			rssItem := dao.RSSItem{}
			rssItem.Title = dataJson.GetString("title")
			rssItem.Link = audioUrl + dataJson.GetString("statistic.sid")
			rssItem.Author = dataJson.GetString("author")
			intro := dataJson.GetString("intro")
			cover := dataJson.GetString("cover")
			rssItem.Content = feed.GenerateContent(intro)
			rssItem.Thumbnail = cover
			items = append(items, rssItem)
		}

		rssData.Items = items
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
