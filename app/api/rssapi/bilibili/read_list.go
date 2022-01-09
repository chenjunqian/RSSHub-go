package bilibili

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
)

func (ctl *Controller) GetReadList(req *ghttp.Request) {
	id := req.GetString("id")

	apiUrl := fmt.Sprintf("https://api.bilibili.com/x/article/list/web/articles?id=%s&jsonp=jsonp", id)
	header := getHeaders()
	header["Referer"] = "https://www.bilibili.com/read/readlist/rl" + id
	rssData := dao.RSSFeed{}
	if resp, err := component.GetHttpClient().SetHeaderMap(header).Get(apiUrl); err == nil {
		respData := gjson.New(resp.ReadAllString())
		dataJson := respData.GetJson("data")
		rssData.Title = "bilibili 专栏文集 - " + dataJson.GetString("list.name")
		rssData.Link = "https://www.bilibili.com/read/readlist/rl" + id
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

		if dataJson.GetString("list.summary") == "" {
			rssData.Description = "作者很懒，还木有写简介.....((/- -)/"
		} else {
			rssData.Description = dataJson.GetString("list.summary")
		}

		articleJsonList := dataJson.GetJsons("articles")
		rssItems := make([]dao.RSSItem, 0)
		for _, articleJson := range articleJsonList {
			rssItem := dao.RSSItem{}
			rssItem.Title = articleJson.GetString("title")
			rssItem.Author = articleJson.GetString("author.name")
			rssItem.Description = feed.GenerateDescription(articleJson.GetString("summary"))
			rssItem.Created = articleJson.GetString("publish_time")
			rssItem.Link = fmt.Sprintf("https://www.bilibili.com/read/cv%s/?from=readlist", articleJson.GetString("id"))
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	_ = req.Response.WriteXmlExit(rssStr)
}
