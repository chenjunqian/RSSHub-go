package bilibili

import (
	"context"
	"fmt"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetUserArticle(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	userId := req.Get("id").String()
	username := getUsernameFromUserId(ctx,userId)

	apiUrl := fmt.Sprintf("https://api.bilibili.com/x/space/article?mid=%s&pn=1&ps=10&sort=publish_time&jsonp=jsonp", userId)
	header := getHeaders()
	header["Referer"] = "https://space.bilibili.com/" + userId
	if resp := service.GetContent(ctx,apiUrl); resp != "" {
		jsonResp := gjson.New(resp)
		articleJsons := jsonResp.GetJsons("data.articles")

		rssData := dao.RSSFeed{}
		rssData.Title = username + " 的 bilibili 专栏"
		rssData.Link = fmt.Sprintf("https://space.bilibili.com/%s/#/article", userId)
		rssData.Description = rssData.Title
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"
		items := make([]dao.RSSItem, 0)
		for _, articleJson := range articleJsons {
			rssItem := dao.RSSItem{}
			rssItem.Title = articleJson.Get("title").String()
			rssItem.Content = feed.GenerateContent(articleJson.Get("summary").String())
			timeStamp := articleJson.Get("publish_time").Int64()
			rssItem.Created = time.Unix(timeStamp, 0).String()
			rssItem.Link = "https://www.bilibili.com/read/cv" + articleJson.Get("id").String()
			items = append(items, rssItem)
		}

		rssData.Items = items
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		req.Response.WriteXmlExit(rssStr)
	}
}
