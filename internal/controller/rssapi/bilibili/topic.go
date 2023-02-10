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

func (ctl *Controller) GetTopic(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	topicName := req.Get("topicName").String()
	apiUrl := "https://api.vc.bilibili.com/topic_svr/v1/topic_svr/topic_new?topic_name=" + topicName
	header := getHeaders()
	header["Referer"] = fmt.Sprintf("https://www.bilibili.com/tag/%s/feed", topicName)
	rssData := dao.RSSFeed{}
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		respJson := gjson.New(resp)
		cardJsonList := respJson.GetJsons("data.cards")

		rssData.Title = topicName + "的全部话题"
		rssData.Link = fmt.Sprintf("https://www.bilibili.com/tag/%s/feed", topicName)
		rssData.Description = rssData.Title
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"
		rssItems := make([]dao.RSSItem, 0)
		for _, cardJson := range cardJsonList {
			cardStr := cardJson.Get("card")
			cardStrJson := gjson.New(cardStr)

			rssItem := dao.RSSItem{}
			var itemTitle string
			if cardStrJson.Get("title") != nil {
				itemTitle = cardStrJson.Get("title").String()
			} else if cardStrJson.Get("description") != nil {
				itemTitle = cardStrJson.Get("description").String()
			} else if cardStrJson.Get("content") != nil {
				itemTitle = cardStrJson.Get("content").String()
			} else if cardStrJson.Get("vest.content") != nil {
				itemTitle = cardStrJson.Get("vest.content").String()
			}
			rssItem.Title = itemTitle

			var imgHtml string
			if picListJson := cardStrJson.GetJsons("pictures"); picListJson != nil {
				for _, picJson := range picListJson {
					imgHtml += fmt.Sprintf("<img src=\"%s\">", picJson.Get("img_src"))
				}
			} else if picJson := cardStrJson.Get("pic").String(); picJson != "" {
				imgHtml += fmt.Sprintf("<img src=\"%s\">", picJson)
			}
			rssItem.Description = rssItem.Title + imgHtml

			if dynamicIdStr := cardJson.Get("desc.dynamic_id_str").String(); dynamicIdStr != "" {
				rssItem.Link = "https://t.bilibili.com/" + dynamicIdStr
			}

			rssItem.Created = cardJson.Get("desc.timestamp").String()
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	req.Response.WriteXmlExit(rssStr)
}
