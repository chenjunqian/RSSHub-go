package bilibili

import (
	"fmt"
	"rsshub/app/dao"
	"rsshub/lib"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
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
	if menuResp, err := g.Client().SetHeaderMap(header).SetCookieMap(getCookieMap()).Get(apiMenuUrl); err == nil {
		menuJsonResp := gjson.New(menuResp.ReadAllString())
		dataJson := menuJsonResp.GetJson("data")
		rssData.Description = dataJson.GetString("intro")
		rssData.Title = dataJson.GetString("title")
	}

	apiUrl := fmt.Sprintf("https://www.bilibili.com/audio/music-service-c/web/song/of-menu?sid=%s&pn=1&ps=100", userId)
	if resp, err := g.Client().SetHeaderMap(header).Get(apiUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		dataJsons := jsonResp.GetJsons("data.data")

		items := make([]dao.RSSItem, 0)
		for _, dataJson := range dataJsons {
			rssItem := dao.RSSItem{}
			rssItem.Title = dataJson.GetString("title")
			rssItem.Link = audioUrl + dataJson.GetString("statistic.sid")
			rssItem.Author = dataJson.GetString("author")
			intro := dataJson.GetString("intro")
			cover := dataJson.GetString("cover")
			rssItem.Description = fmt.Sprintf("%s<br><img src='%s'>", intro, cover)
			items = append(items, rssItem)
		}

		rssData.Items = items
		rssStr := lib.GenerateRSS(rssData)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
