package bilibili

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) GetUserFav(req *ghttp.Request) {
	id := req.GetString("id")
	username := getUsernameFromUserId(id)
	apiUrl := fmt.Sprintf("https://api.bilibili.com/x/v2/fav/video?vmid=%s&ps=30&tid=0&keyword=&pn=1&order=fav_time", id)
	header := getHeaders()
	header["Referer"] = fmt.Sprintf("https://space.bilibili.com/%s/#/favlist", id)
	rssData := dao.RSSFeed{}
	if resp, err := g.Client().SetHeaderMap(header).Get(apiUrl); err == nil {
		dataJson := gjson.New(resp.ReadAllString())
		rssData.Title = username + " 的 bilibili 收藏夹"
		rssData.Link = fmt.Sprintf("https://space.bilibili.com/%s/#/favlist", id)
		rssData.Description = rssData.Title
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

		archiveJsonList := dataJson.GetJsons("data.archives")
		rssItems := make([]dao.RSSItem, 0)
		for _, archiveJson := range archiveJsonList {
			rssItem := dao.RSSItem{}
			rssItem.Title = archiveJson.GetString("title")
			desc := archiveJson.GetString("desc")
			picId := archiveJson.GetString("pic")
			rssItem.Description = fmt.Sprintf("<img src=\"%s\"><br>%s", picId, desc)
			rssItem.Created = archiveJson.GetString("fav_at")
			if archiveJson.GetFloat64("fav_at") > archiveJson.GetFloat64("bvidTime") && archiveJson.GetFloat64("fav_at") > archiveJson.GetFloat64("bvid") {
				rssItem.Link = archiveJson.GetString("bvid")
			} else {
				rssItem.Link = archiveJson.GetString("aid")
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := lib.GenerateRSS(rssData)
	_ = req.Response.WriteXmlExit(rssStr)
}
