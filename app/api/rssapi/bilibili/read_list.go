package bilibili

import (
	"context"
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetReadList(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	id := req.Get("id").String()

	apiUrl := fmt.Sprintf("https://api.bilibili.com/x/article/list/web/articles?id=%s&jsonp=jsonp", id)
	header := getHeaders()
	header["Referer"] = "https://www.bilibili.com/read/readlist/rl" + id
	rssData := dao.RSSFeed{}
	if resp, err := component.GetHttpClient().SetHeaderMap(header).Get(ctx, apiUrl); err == nil {
		defer resp.Close()
		respData := gjson.New(resp.ReadAllString())
		dataJson := respData.GetJson("data")
		rssData.Title = "bilibili 专栏文集 - " + dataJson.Get("list.name").String()
		rssData.Link = "https://www.bilibili.com/read/readlist/rl" + id
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

		if dataJson.Get("list.summary").String() == "" {
			rssData.Description = "作者很懒，还木有写简介.....((/- -)/"
		} else {
			rssData.Description = dataJson.Get("list.summary").String()
		}

		articleJsonList := dataJson.GetJsons("articles")
		rssItems := make([]dao.RSSItem, 0)
		for _, articleJson := range articleJsonList {
			rssItem := dao.RSSItem{}
			rssItem.Title = articleJson.Get("title").String()
			rssItem.Author = articleJson.Get("author.name").String()
			rssItem.Content = feed.GenerateContent(articleJson.Get("summary").String())
			rssItem.Created = articleJson.Get("publish_time").String()
			rssItem.Link = fmt.Sprintf("https://www.bilibili.com/read/cv%s/?from=readlist", articleJson.Get("id"))
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	req.Response.WriteXmlExit(rssStr)
}
