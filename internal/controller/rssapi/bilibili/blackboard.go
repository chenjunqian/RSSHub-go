package bilibili

import (
	"context"
	"fmt"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetBlackboard(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	apiUrl := "https://www.bilibili.com/activity/page/list?plat=1,2,3&mold=1&http=3&page=1&tid=0"
	header := getHeaders()
	header["Referer"] = "https://www.bilibili.com/blackboard/topic_list.html"

	rssData := dao.RSSFeed{}
	rssData.Title = "bilibili 话题列表"
	rssData.Link = "https://www.bilibili.com/blackboard/topic_list.html#/"
	rssData.Description = "bilibili 话题列表"
	rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"
	if resp := service.GetContent(ctx,apiUrl); resp != "" {
		jsonResp := gjson.New(resp)
		dataJsonList := jsonResp.GetJsons("data.list")
		items := make([]dao.RSSItem, 0)
		nameList := garray.New()
		for _, dataJson := range dataJsonList {
			title := dataJson.Get("name").String()
			if nameList.Contains(title) {
				continue
			} else {
				nameList.Append(title)
			}
			rssItem := dao.RSSItem{}
			rssItem.Title = title
			rssItem.Content = fmt.Sprintf("%s<br> %s", dataJson.Get("name"), dataJson.Get("desc"))
			rssItem.Link = dataJson.Get("pc_url").String()
			rssItem.Created = dataJson.Get("ctime").String()
			items = append(items, rssItem)
		}
		rssData.Items = items
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	req.Response.WriteXmlExit(rssStr)
}
