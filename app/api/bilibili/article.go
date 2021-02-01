package bilibili

import (
	"fmt"
	"rsshub/app/dao"
	"rsshub/lib"
	"time"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetUserArticle(req *ghttp.Request) {
	userId := req.GetString("id")
	username := getUsernameFromUserId(userId)

	apiUrl := fmt.Sprintf("https://api.bilibili.com/x/space/article?mid=%s&pn=1&ps=10&sort=publish_time&jsonp=jsonp", userId)
	header := getHeaders()
	header["Referer"] = "https://space.bilibili.com/" + userId
	if resp, err := g.Client().SetHeaderMap(header).Get(apiUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		articleJsons := jsonResp.GetJsons("data.articles")

		rssData := dao.RSSFeed{}
		rssData.Title = username + " 的 bilibili 专栏"
		rssData.Link = fmt.Sprintf("https://space.bilibili.com/%s/#/article", userId)
		rssData.Description = rssData.Title
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"
		items := make([]dao.RSSItem, 0)
		for _, articleJson := range articleJsons {
			rssItem := dao.RSSItem{}
			rssItem.Title = articleJson.GetString("title")
			rssItem.Description = lib.GenerateDescription(articleJson.GetString("image_urls.0"), articleJson.GetString("summary"))
			timeStamp := articleJson.GetInt64("publish_time")
			rssItem.Created = time.Unix(timeStamp, 0).String()
			rssItem.Link = "https://www.bilibili.com/read/cv" + articleJson.GetString("id")
			items = append(items, rssItem)
		}

		rssData.Items = items
		rssStr := lib.GenerateRSS(rssData)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
