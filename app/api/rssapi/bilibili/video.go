package bilibili

import (
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetUserVideo(req *ghttp.Request) {
	id := req.GetString("id")
	username := getUsernameFromUserId(id)
	apiUrl := fmt.Sprintf("https://api.bilibili.com/x/space/arc/search?mid=%s&ps=10&tid=0&pn=1&order=pubdate&jsonp=jsonp", id)
	header := getHeaders()
	header["Referer"] = fmt.Sprintf("https://space.bilibili.com/%s/", id)
	rssData := dao.RSSFeed{}
	if resp := component.GetContent(apiUrl); resp != "" {
		dataJson := gjson.New(resp)
		rssData.Title = username + " 的 bilibili 空间"
		rssData.Link = "https://space.bilibili.com/" + id
		rssData.Description = rssData.Title
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

		videoJsonList := dataJson.GetJsons("data.list.vlist")
		rssItems := make([]dao.RSSItem, 0)
		for _, videoJson := range videoJsonList {
			rssItem := dao.RSSItem{}
			rssItem.Title = videoJson.GetString("title")
			desc := videoJson.GetString("description")
			picId := videoJson.GetString("pic")
			rssItem.Description = fmt.Sprintf("<img src=\"%s\"><br>%s", picId, desc)
			rssItem.Created = videoJson.GetString("created")
			if videoJson.GetFloat64("created") > videoJson.GetFloat64("bvidTime") && videoJson.GetFloat64("created") > videoJson.GetFloat64("bvid") {
				rssItem.Link = "https://www.bilibili.com/video/" + videoJson.GetString("bvid")
			} else {
				rssItem.Link = "https://www.bilibili.com/video/av" + videoJson.GetString("aid")
			}
			rssItem.Author = username
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	_ = req.Response.WriteXmlExit(rssStr)
}
