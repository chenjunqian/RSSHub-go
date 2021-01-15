package bilibili

import (
	"fmt"
	"rsshub/app/dao"
	"rsshub/lib"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetReadList(req *ghttp.Request) {
	id := req.GetString("id")

	apiUrl := fmt.Sprintf("https://api.bilibili.com/x/article/list/web/articles?id=%s&jsonp=jsonp", id)
	header := getHeaders()
	header["Referer"] = "https://www.bilibili.com/read/readlist/rl" + id
	rssData := dao.RSSFeed{}
	if resp, err := g.Client().SetHeaderMap(header).Get(apiUrl); err == nil {
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
			rssItem.Description = fmt.Sprintf("%s…<br><img src='%s'>", articleJson.GetString("summary"), articleJson.GetString("image_urls.0"))
			rssItem.Created = articleJson.GetString("publish_time")
			rssItem.Link = fmt.Sprintf("https://www.bilibili.com/read/cv%s/?from=readlist", articleJson.GetString("id"))
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData)
	_ = req.Response.WriteXmlExit(rssStr)
}
