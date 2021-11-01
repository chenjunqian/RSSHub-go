package bilibili

import (
	"fmt"
	"rsshub/app/dao"
	"rsshub/lib"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetBlackboard(req *ghttp.Request) {
	apiUrl := "https://www.bilibili.com/activity/page/list?plat=1,2,3&mold=1&http=3&page=1&tid=0"
	header := getHeaders()
	header["Referer"] = "https://www.bilibili.com/blackboard/topic_list.html"

	rssData := dao.RSSFeed{}
	rssData.Title = "bilibili 话题列表"
	rssData.Link = "https://www.bilibili.com/blackboard/topic_list.html#/"
	rssData.Description = "bilibili 话题列表"
	rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"
	if resp, err := g.Client().SetHeaderMap(header).Get(apiUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		dataJsonList := jsonResp.GetJsons("data.list")
		items := make([]dao.RSSItem, 0)
		nameList := garray.New()
		for _, dataJson := range dataJsonList {
			title := dataJson.GetString("name")
			if nameList.Contains(title) {
				continue
			} else {
				nameList.Append(title)
			}
			rssItem := dao.RSSItem{}
			rssItem.Title = title
			rssItem.Description = fmt.Sprintf("%s<br> %s", dataJson.GetString("name"), dataJson.GetString("desc"))
			rssItem.Link = dataJson.GetString("pc_url")
			rssItem.Created = dataJson.GetString("ctime")
			items = append(items, rssItem)
		}
		rssData.Items = items
	}

	rssStr := lib.GenerateRSS(rssData, req.Router.Uri)
	_ = req.Response.WriteXmlExit(rssStr)
}
