package bilibili

import (
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetLinkNews(req *ghttp.Request) {
	product := req.GetString("product")

	var productTitle string
	switch product {
	case "live":
		productTitle = "直播"
		break
	case "vc":
		productTitle = "小视频"
		break
	case "wh":
		productTitle = "相簿"
		break
	}

	rssData := dao.RSSFeed{}
	apiUrl := fmt.Sprintf("https://api.vc.bilibili.com/news/v1/notice/list?platform=pc&product=%s&category=all&page_no=1&page_size=20", product)
	header := getHeaders()
	header["Referer"] = "https://link.bilibili.com/p/eden/news"
	if resp := component.GetContent(apiUrl); resp != "" {
		jsonResp := gjson.New(resp)
		rssData.Title = fmt.Sprintf("bilibili %s公告", productTitle)
		rssData.Link = fmt.Sprintf("https://link.bilibili.com/p/eden/news#/?tab=%s&tag=all&page_no=1", product)
		rssData.Description = rssData.Title
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

		items := make([]dao.RSSItem, 0)
		itemJsons := jsonResp.GetJsons("data.items")
		for _, item := range itemJsons {
			rssItem := dao.RSSItem{}
			rssItem.Title = item.GetString("title")
			rssItem.Content = feed.GenerateContent(item.GetString("mark"))
			rssItem.Created = item.GetString("ctime")
			if item.GetString("announce_link") != "" {
				rssItem.Link = item.GetString("announce_link")
			} else {
				rssItem.Link = "https://link.bilibili.com/p/eden/news#/newsdetail?id=" + item.GetString("id")
			}
			items = append(items, rssItem)
		}

		rssData.Items = items
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	_ = req.Response.WriteXmlExit(rssStr)
}
