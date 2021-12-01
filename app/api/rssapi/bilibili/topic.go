package bilibili

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
)

func (ctl *Controller) GetTopic(req *ghttp.Request) {
	topicName := req.GetString("topicName")
	apiUrl := "https://api.vc.bilibili.com/topic_svr/v1/topic_svr/topic_new?topic_name=" + topicName
	header := getHeaders()
	header["Referer"] = fmt.Sprintf("https://www.bilibili.com/tag/%s/feed", topicName)
	rssData := dao.RSSFeed{}
	if resp, err := g.Client().SetHeaderMap(header).Get(apiUrl); err == nil {
		respJson := gjson.New(resp.ReadAllString())
		cardJsonList := respJson.GetJsons("data.cards")

		rssData.Title = topicName + "的全部话题"
		rssData.Link = fmt.Sprintf("https://www.bilibili.com/tag/%s/feed", topicName)
		rssData.Description = rssData.Title
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"
		rssItems := make([]dao.RSSItem, 0)
		for _, cardJson := range cardJsonList {
			cardStr := cardJson.GetString("card")
			cardStrJson := gjson.New(cardStr)

			rssItem := dao.RSSItem{}
			var itemTitle string
			if cardStrJson.Get("title") != nil {
				itemTitle = cardStrJson.GetString("title")
			} else if cardStrJson.Get("description") != nil {
				itemTitle = cardStrJson.GetString("description")
			} else if cardStrJson.Get("content") != nil {
				itemTitle = cardStrJson.GetString("content")
			} else if cardStrJson.Get("vest.content") != nil {
				itemTitle = cardStrJson.GetString("vest.content")
			}
			rssItem.Title = itemTitle

			var imgHtml string
			if picListJson := cardStrJson.GetJsons("pictures"); picListJson != nil {
				for _, picJson := range picListJson {
					imgHtml += fmt.Sprintf("<img src=\"%s\">", picJson.GetString("img_src"))
				}
			} else if picJson := cardStrJson.GetString("pic"); picJson != "" {
				imgHtml += fmt.Sprintf("<img src=\"%s\">", picJson)
			}
			rssItem.Description = rssItem.Title + imgHtml

			if dynamicIdStr := cardJson.GetString("desc.dynamic_id_str"); dynamicIdStr != "" {
				rssItem.Link = "https://t.bilibili.com/" + dynamicIdStr
			}

			rssItem.Created = cardJson.GetString("desc.timestamp")
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	_ = req.Response.WriteXmlExit(rssStr)
}
