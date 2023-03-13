package zhihu

import (
	"context"
	"fmt"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetCollections(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	colletionId := req.Get("id")
	collectionGetUrl := fmt.Sprintf("https://www.zhihu.com/collection/%s", colletionId)
	headers := getHeaders()
	headers["Referer"] = fmt.Sprintf("https://www.zhihu.com/people/%s/activities", colletionId)
	cookieMap := getCookieMap(ctx)
	if resp, err := service.GetHttpClient().SetHeaderMap(headers).SetCookieMap(cookieMap).Get(ctx, collectionGetUrl); err == nil {
		doc := soup.HTMLParse(resp.ReadAllString())

		rssData := dao.RSSFeed{}
		collectionTitle := doc.Find("div", "class", "CollectionDetailPageHeader-title").Text()
		rssData.Title = collectionTitle
		rssData.Link = collectionGetUrl
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"

		items := make([]dao.RSSItem, 0)
		collectionList := doc.FindAll("div", "class", "CollectionDetailPageItem-innerContainer")
		fmt.Println("collectionList length : ", len(collectionList))
		for _, collectionItem := range collectionList {
			titleTag := collectionItem.Find("h2", "class", "ContentItem-title").Find("a")
			itemTitle := titleTag.Text()
			linkUrl := titleTag.Attrs()["href"]
			image := collectionItem.Find("meta", "itemprop", "image").Attrs()["content"]
			text := collectionItem.Find("span", "class", "CopyrightRichText-richText").Text()
			var description string
			if image != "" {
				description = fmt.Sprintf("%s <img src='%s'>", text, image)
			} else {
				description = text
			}

			rssItem := dao.RSSItem{
				Title:       itemTitle,
				Description: description,
				Link:        linkUrl,
			}
			items = append(items, rssItem)
		}

		rssData.Items = items
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		req.Response.WriteXmlExit(rssStr)
	}
}
