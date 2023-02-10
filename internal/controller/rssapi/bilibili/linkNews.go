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

func (ctl *Controller) GetLinkNews(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	product := req.Get("product").String()

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
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		jsonResp := gjson.New(resp)
		rssData.Title = fmt.Sprintf("bilibili %s公告", productTitle)
		rssData.Link = fmt.Sprintf("https://link.bilibili.com/p/eden/news#/?tab=%s&tag=all&page_no=1", product)
		rssData.Description = rssData.Title
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

		items := make([]dao.RSSItem, 0)
		itemJsons := jsonResp.GetJsons("data.items")
		for _, item := range itemJsons {
			rssItem := dao.RSSItem{}
			rssItem.Title = item.Get("title").String()
			rssItem.Content = feed.GenerateContent(item.Get("mark").String())
			rssItem.Created = item.Get("ctime").String()
			if item.Get("announce_link").String() != "" {
				rssItem.Link = item.Get("announce_link").String()
			} else {
				rssItem.Link = "https://link.bilibili.com/p/eden/news#/newsdetail?id=" + item.Get("id").String()
			}
			items = append(items, rssItem)
		}

		rssData.Items = items
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	req.Response.WriteXmlExit(rssStr)
}
