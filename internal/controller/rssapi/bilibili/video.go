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

func (ctl *Controller) GetUserVideo(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	id := req.Get("id").String()
	username := getUsernameFromUserId(ctx, id)
	apiUrl := fmt.Sprintf("https://api.bilibili.com/x/space/arc/search?mid=%s&ps=10&tid=0&pn=1&order=pubdate&jsonp=jsonp", id)
	header := getHeaders()
	header["Referer"] = fmt.Sprintf("https://space.bilibili.com/%s/", id)
	rssData := dao.RSSFeed{}
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		dataJson := gjson.New(resp)
		rssData.Title = username + " 的 bilibili 空间"
		rssData.Link = "https://space.bilibili.com/" + id
		rssData.Description = rssData.Title
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

		videoJsonList := dataJson.GetJsons("data.list.vlist")
		rssItems := make([]dao.RSSItem, 0)
		for _, videoJson := range videoJsonList {
			rssItem := dao.RSSItem{}
			rssItem.Title = videoJson.Get("title").String()
			desc := videoJson.Get("description")
			picId := videoJson.Get("pic")
			rssItem.Description = fmt.Sprintf("<img src=\"%s\"><br>%s", picId, desc)
			rssItem.Created = videoJson.Get("created").String()
			if videoJson.Get("created").Float64() > videoJson.Get("bvidTime").Float64() && videoJson.Get("created").Float64() > videoJson.Get("bvid").Float64() {
				rssItem.Link = "https://www.bilibili.com/video/" + videoJson.Get("bvid").String()
			} else {
				rssItem.Link = "https://www.bilibili.com/video/av" + videoJson.Get("aid").String()
			}
			rssItem.Author = username
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	req.Response.WriteXmlExit(rssStr)
}
