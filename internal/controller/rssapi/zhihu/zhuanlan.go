package zhihu

import (
	"context"
	"fmt"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetZhuanlan(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	zhuanlanId := req.Get("id")
	zhuanlanUrl := fmt.Sprintf("https://www.zhihu.com/api/v4/columns/%s/items", zhuanlanId)
	headers := getHeaders()
	headers["Referer"] = fmt.Sprintf("https://zhuanlan.zhihu.com/%s", zhuanlanId)
	if resp, err := service.GetHttpClient().SetHeaderMap(headers).Get(ctx, zhuanlanUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.Get("data").Strings()

		infoUrl := fmt.Sprintf("https://zhuanlan.zhihu.com/%s", zhuanlanId)
		infoResp, _ := service.GetHttpClient().SetHeaderMap(headers).Get(ctx, infoUrl)
		doc := soup.HTMLParse(infoResp.ReadAllString())

		feedTitle := doc.Find("div", "class", "css-zyehvu").Text()
		feedLink := doc.Find("div", "class", "css-1bnklpv").Text()
		rssData := dao.RSSFeed{}
		rssData.Title = feedTitle
		rssData.Link = feedLink
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"

		items := make([]dao.RSSItem, 0)
		for index := range respDataList {
			content := jsonResp.Get(fmt.Sprintf("data.%d.content", index)).String()
			contentType := jsonResp.Get(fmt.Sprintf("data.%d.type", index)).String()

			var itemTitle string
			var itemLink string
			var itemAuthor string
			var itemCreated string
			switch contentType {
			case "answer":
				itemTitle = jsonResp.Get(fmt.Sprintf("data.%d.question.title", index)).String()
				itemAuthor = jsonResp.Get(fmt.Sprintf("data.%d.question.author.name", index)).String()
				questionId := jsonResp.Get(fmt.Sprintf("data.%d.question.id", index))
				answerId := jsonResp.Get(fmt.Sprintf("data.%d.id", index))
				itemLink = fmt.Sprintf("https://www.zhihu.com/question/%s/answer/%s", questionId, answerId)
				itemCreated = jsonResp.Get(fmt.Sprintf("data.%d.created", index)).String()
			case "article":
				itemTitle = jsonResp.Get(fmt.Sprintf("data.%d.title", index)).String()
				itemAuthor = jsonResp.Get(fmt.Sprintf("data.%d.author.name", index)).String()
				itemLink = jsonResp.Get(fmt.Sprintf("data.%d.link", index)).String()
				itemCreated = jsonResp.Get(fmt.Sprintf("data.%d.created", index)).String()
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
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		req.Response.WriteXmlExit(rssStr)
	}
}
