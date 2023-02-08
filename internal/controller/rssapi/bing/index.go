package bing

import (
	"context"
	"fmt"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetDailyImage(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	if value, err := cache.GetCache(ctx, "BING_DAILY_IMG"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	baseUrl := "http://www.bing.com"
	apiUrl := baseUrl + "/HPImageArchive.aspx?format=js&idx=0&n=1"
	rssData := dao.RSSFeed{}
	rssData.Title = "Bing每日壁纸"
	rssData.Link = "https://cn.bing.com/"
	rssData.Tag = []string{"壁纸", "图片"}
	rssData.ImageUrl = "https://cn.bing.com/sa/simg/favicon-2x.ico"
	if statusResp := service.GetContent(ctx,apiUrl); statusResp != "" {
		respJson := gjson.New(statusResp)
		imageListJson := respJson.GetJsons("images")
		rssItems := make([]dao.RSSItem, 0)
		for _, imageJson := range imageListJson {
			rssItem := dao.RSSItem{
				Title: imageJson.Get("copyright").String(),
				Link:  imageJson.Get("copyrightlink").String(),
			}
			var content string
			content = content + fmt.Sprintf("<img src=\"%s%s\">", baseUrl, imageJson.Get("url"))
			rssItem.Description = content
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx,"BING_DAILY_IMG", rssStr)
	req.Response.WriteXmlExit(rssStr)
}
