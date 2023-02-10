package mitchina

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetFlash(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	if value, err := cache.GetCache(ctx, "MIT_CHINA_FLASH"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "http://i.mittrchina.com/flash"
	rssData := dao.RSSFeed{
		Title:       "MIT 中国",
		Link:        "http://www.mittrchina.com/newsflash",
		Description: "MIT科技评论快讯",
		Tag:         []string{"科技"},
		ImageUrl:    "https://www.mittrchina.com/logo.ico",
	}
	if resp, err := service.GetHttpClient().SetHeaderMap(getHeaders()).Post(ctx, apiUrl, map[string]string{"page": "1", "size": "10"}); err == nil {
		defer func(resp *gclient.Response) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}(resp)
		respJson := gjson.New(resp.ReadAllString())
		itemJsonList := respJson.GetJsons("data.items")
		rssItems := make([]dao.RSSItem, 0)
		for _, itemJson := range itemJsonList {
			time := itemJson.Get("push_time").String()
			title := itemJson.Get("name").String()
			content := itemJson.Get("content").String()
			rssItem := dao.RSSItem{
				Title:   title,
				Link:    "http://www.mittrchina.com/newsflash",
				Content: feed.GenerateContent(content),
				Author:  "MIT 中国",
				Created: time,
			}

			rssItems = append(rssItems, rssItem)
		}

		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx, "MIT_CHINA_FLASH", rssStr)
	req.Response.WriteXmlExit(rssStr)
}
