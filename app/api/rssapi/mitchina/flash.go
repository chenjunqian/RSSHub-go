package mitchina

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) GetFlash(req *ghttp.Request) {
	if value, err := g.Redis().DoVar("GET", "MIT_CHINA_FLASH"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "http://i.mittrchina.com/flash"
	rssData := dao.RSSFeed{
		Title:       "MIT 中国",
		Link:        "http://www.mittrchina.com/newsflash",
		Description: "MIT科技评论快讯",
		Tag:         []string{"科技"},
		ImageUrl:    "http://www.mittrchina.com/logo.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Post(apiUrl, map[string]string{"page": "1", "size": "10"}); err == nil {
		respJson := gjson.New(resp.ReadAllString())
		itemJsonList := respJson.GetJsons("data.items")
		rssItems := make([]dao.RSSItem, 0)
		for _, itemJson := range itemJsonList {
			time := itemJson.GetString("push_time")
			title := itemJson.GetString("name")
			content := itemJson.GetString("content")
			rssItem := dao.RSSItem{
				Title:       title,
				Link:        "http://www.mittrchina.com/newsflash",
				Description: lib.GenerateDescription("", content),
				Author:      "MIT 中国",
				Created:     time,
			}

			rssItems = append(rssItems, rssItem)
		}

		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", "MIT_CHINA_FLASH", rssStr)
	g.Redis().DoVar("EXPIRE", "MIT_CHINA_FLASH", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
