package bilibili

import (
	"fmt"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strconv"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetLinkRoom(req *ghttp.Request) {
	roomId := req.GetString("roomId")
	roomIdInt, _ := strconv.ParseInt(roomId, 10, 64)
	if roomIdInt < 10000 {
		roomId = getLiveIDFromShortID(roomId)
	}

	name := getUsernameFromLiveID(roomId)
	rssData := dao.RSSFeed{}
	rssData.Title = fmt.Sprintf("%s 直播间开播状态", name)
	rssData.Link = "https://live.bilibili.com/" + roomId
	rssData.Description = rssData.Title
	rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

	apiUrl := fmt.Sprintf("https://api.live.bilibili.com/room/v1/Room/get_info?room_id=%s&from=room", roomId)
	header := getHeaders()
	header["Referer"] = "https://live.bilibili.com/" + roomId
	if resp, err := g.Client().SetHeaderMap(header).Get(apiUrl); err == nil {
		respJson := gjson.New(resp.ReadAllString())
		dataJson := respJson.GetJson("data")

		rssItem := dao.RSSItem{}
		if dataJson.GetInt64("live_status") == 1 {
			rssItem.Title = fmt.Sprintf("%s %s", dataJson.GetString("title"), dataJson.GetString("live_time"))
			rssItem.Description = fmt.Sprintf("%s<br>%s", dataJson.GetString("title"), dataJson.GetString("description"))
			rssItem.Created = dataJson.GetString("live_time")
			rssItem.Link = "https://live.bilibili.com/" + roomId
			rssData.Items = []dao.RSSItem{rssItem}
		}
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	_ = req.Response.WriteXmlExit(rssStr)
}
