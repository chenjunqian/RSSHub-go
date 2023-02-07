package idaily

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"
	"strconv"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {

	var ctx context.Context = context.Background()
	if value, err := service.GetRedis().Do(ctx,"GET", "IDAILY_INDEX"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://idaily-cdn.idailycdn.com/api/list/v3/iphone/zh-hans?page=1&ver=iphone"
	rssData := dao.RSSFeed{
		Title:       "iDaily · 每日环球视野",
		Link:        apiUrl,
		Tag:         []string{"海外"},
		Description: "iDaily 每日环球视野",
		ImageUrl:    "https://dayoneapp.com/favicon-32x32.png?v=9277df7ae7503b6e383587ae0e7210ee",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != "" {
		respJson := gjson.New(resp)
		dataJsonArray := respJson.Array()
		rssItems := make([]dao.RSSItem, 0)
		for index := range dataJsonArray {
			var (
				title          string
				link           string
				time           string
				content        string
				coverImageLink string
				dataJson       *gjson.Json
			)
			dataJson = respJson.GetJson(strconv.Itoa(index))
			link = dataJson.Get("link_share").String()
			time = dataJson.Get("pubdate_timestamp").String()
			title = dataJson.Get("ui_sets.caption_subtitle").String()
			coverImageLink = dataJson.Get("cover_landscape").String()
			content = dataJson.Get("content").String()
			rssItem := dao.RSSItem{
				Title:     title,
				Link:      link,
				Content:   feed.GenerateContent(content),
				Created:   time,
				Thumbnail: coverImageLink,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	service.GetRedis().Do(ctx,"SET", "IDAILY_INDEX", rssStr)
	service.GetRedis().Do(ctx,"EXPIRE", "IDAILY_INDEX", 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
