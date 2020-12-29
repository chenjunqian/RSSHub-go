package zhihu

import (
	"fmt"
	"rsshub/lib"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetCollections(req *ghttp.Request) {
	colletionId := req.Get("id")
	collectionGetUrl := fmt.Sprintf("https://www.zhihu.com/collection/%s", colletionId)
	headers := getHeaders()
	headers["Referer"] = fmt.Sprintf("https://www.zhihu.com/people/%s/activities", colletionId)
	cookieMap := getCookieMap()
	if resp, err := g.Client().SetHeaderMap(headers).SetCookieMap(cookieMap).Get(collectionGetUrl); err == nil {
		doc := soup.HTMLParse(resp.ReadAllString())

		rssData := make(map[string]interface{})
		collectionTitle := doc.Find("div", "class", "CollectionDetailPageHeader-title").Text()
		rssData["title"] = collectionTitle
		rssData["link"] = collectionGetUrl

		items := make([]map[string]string, 0)
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

			itemMap := make(map[string]string)
			itemMap["title"] = itemTitle
			itemMap["description"] = description
			itemMap["link"] = linkUrl
			items = append(items, itemMap)
		}

		rssData["items"] = items
		rssStr := lib.GenerateRSS(rssData)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
