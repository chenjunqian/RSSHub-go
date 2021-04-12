package zhihu

import (
	"fmt"
	"rsshub/app/dao"
	"rsshub/lib"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetZhuanlan(req *ghttp.Request) {
	zhuanlanId := req.Get("id")
	zhuanlanUrl := fmt.Sprintf("https://www.zhihu.com/api/v4/columns/%s/items", zhuanlanId)
	headers := getHeaders()
	headers["Referer"] = fmt.Sprintf("https://zhuanlan.zhihu.com/%s", zhuanlanId)
	if resp, err := g.Client().SetHeaderMap(headers).Get(zhuanlanUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.GetArray("data")

		infoUrl := fmt.Sprintf("https://zhuanlan.zhihu.com/%s", zhuanlanId)
		infoResp, _ := g.Client().SetHeaderMap(headers).Get(infoUrl)
		doc := soup.HTMLParse(infoResp.ReadAllString())

		feedTitle := doc.Find("div", "class", "css-zyehvu").Text()
		feedLink := doc.Find("div", "class", "css-1bnklpv").Text()
		rssData := dao.RSSFeed{}
		rssData.Title = feedTitle
		rssData.Link = feedLink
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"

		items := make([]dao.RSSItem, 0)
		for index := range respDataList {
			content := jsonResp.GetString(fmt.Sprintf("data.%d.content", index))
			contentType := jsonResp.GetString(fmt.Sprintf("data.%d.type", index))

			var itemTitle string
			var itemLink string
			var itemAuthor string
			var itemCreated string
			switch contentType {
			case "answer":
				itemTitle = jsonResp.GetString(fmt.Sprintf("data.%d.question.title", index))
				itemAuthor = jsonResp.GetString(fmt.Sprintf("data.%d.question.author.name", index))
				questionId := jsonResp.GetString(fmt.Sprintf("data.%d.question.id", index))
				answerId := jsonResp.GetString(fmt.Sprintf("data.%d.id", index))
				itemLink = fmt.Sprintf("https://www.zhihu.com/question/%s/answer/%s", questionId, answerId)
				itemCreated = jsonResp.GetString(fmt.Sprintf("data.%d.created", index))
			case "article":
				itemTitle = jsonResp.GetString(fmt.Sprintf("data.%d.title", index))
				itemAuthor = jsonResp.GetString(fmt.Sprintf("data.%d.author.name", index))
				itemLink = jsonResp.GetString(fmt.Sprintf("data.%d.link", index))
				itemCreated = jsonResp.GetString(fmt.Sprintf("data.%d.created", index))
			}

			rssItem := dao.RSSItem{
				Title:       itemTitle,
				Description: content,
				Link:        itemLink,
				Author:      itemAuthor,
				Created:     itemCreated,
			}
			items = append(items, rssItem)
		}

		rssData.Items = items
		rssStr := lib.GenerateRSS(rssData)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
