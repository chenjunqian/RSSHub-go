package bilibili

import (
	"fmt"
	"rsshub/app/dao"
	"rsshub/lib"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
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
	if resp, err := g.Client().SetHeaderMap(header).Get(apiUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		rssData.Title = fmt.Sprintf("bilibili %s公告", productTitle)
		rssData.Link = fmt.Sprintf("https://link.bilibili.com/p/eden/news#/?tab=%s&tag=all&page_no=1", product)
		rssData.Description = rssData.Title

		items := make([]dao.RSSItem, 0)
		itemJsons := jsonResp.GetJsons("data.items")
		for _, item := range itemJsons {
			rssItem := dao.RSSItem{}
			rssItem.Title = item.GetString("title")
			rssItem.Description = fmt.Sprintf("%s<br><img src='%s'>", item.GetString("mark"), item.GetString("cover_url"))
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
	rssStr := lib.GenerateRSS(rssData)
	_ = req.Response.WriteXmlExit(rssStr)
}
