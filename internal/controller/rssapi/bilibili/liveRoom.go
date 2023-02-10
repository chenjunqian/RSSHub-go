package bilibili

import (
	"context"
	"fmt"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"
	"strconv"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetLinkRoom(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	roomId := req.Get("roomId").String()
	roomIdInt, _ := strconv.ParseInt(roomId, 10, 64)
	if roomIdInt < 10000 {
		roomId = getLiveIDFromShortID(ctx, roomId)
	}

	name := getUsernameFromLiveID(ctx, roomId)
	rssData := dao.RSSFeed{}
	rssData.Title = fmt.Sprintf("%s 直播间开播状态", name)
	rssData.Link = "https://live.bilibili.com/" + roomId
	rssData.Description = rssData.Title
	rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

	apiUrl := fmt.Sprintf("https://api.live.bilibili.com/room/v1/Room/get_info?room_id=%s&from=room", roomId)
	header := getHeaders()
	header["Referer"] = "https://live.bilibili.com/" + roomId
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		respJson := gjson.New(resp)
		dataJson := respJson.GetJson("data")

		rssItem := dao.RSSItem{}
		if dataJson.Get("live_status").Int64() == 1 {
			rssItem.Title = fmt.Sprintf("%s %s", dataJson.Get("title"), dataJson.Get("live_time"))
			rssItem.Description = fmt.Sprintf("%s<br>%s", dataJson.Get("title"), dataJson.Get("description"))
			rssItem.Created = dataJson.Get("live_time").String()
			rssItem.Link = "https://live.bilibili.com/" + roomId
			rssData.Items = []dao.RSSItem{rssItem}
		}
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	req.Response.WriteXmlExit(rssStr)
}
