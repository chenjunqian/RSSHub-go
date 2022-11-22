package bilibili

import (
	"context"
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetUserFav(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	id := req.Get("id").String()
	username := getUsernameFromUserId(ctx,id)
	apiUrl := fmt.Sprintf("https://api.bilibili.com/x/v2/fav/video?vmid=%s&ps=30&tid=0&keyword=&pn=1&order=fav_time", id)
	header := getHeaders()
	header["Referer"] = fmt.Sprintf("https://space.bilibili.com/%s/#/favlist", id)
	rssData := dao.RSSFeed{}
	if resp := component.GetContent(ctx,apiUrl); resp != "" {
		dataJson := gjson.New(resp)
		rssData.Title = username + " 的 bilibili 收藏夹"
		rssData.Link = fmt.Sprintf("https://space.bilibili.com/%s/#/favlist", id)
		rssData.Description = rssData.Title
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

		archiveJsonList := dataJson.GetJsons("data.archives")
		rssItems := make([]dao.RSSItem, 0)
		for _, archiveJson := range archiveJsonList {
			rssItem := dao.RSSItem{}
			rssItem.Title = archiveJson.Get("title").String()
			desc := archiveJson.Get("desc")
			picId := archiveJson.Get("pic")
			rssItem.Description = fmt.Sprintf("<img src=\"%s\"><br>%s", picId, desc)
			rssItem.Created = archiveJson.Get("fav_at").String()
			if archiveJson.Get("fav_at").Float64() > archiveJson.Get("bvidTime").Float64() && archiveJson.Get("fav_at").Float64() > archiveJson.Get("bvid").Float64() {
				rssItem.Link = archiveJson.Get("bvid").String()
			} else {
				rssItem.Link = archiveJson.Get("aid").String()
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	req.Response.WriteXmlExit(rssStr)
}
