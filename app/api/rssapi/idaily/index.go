package idaily

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strconv"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {

	if value, err := g.Redis().DoVar("GET", "IDAILY_INDEX"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
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
	if resp, err := component.GetHttpClient().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		respJson := gjson.New(resp.ReadAllString())
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
			link = dataJson.GetString("link_share")
			time = dataJson.GetString("pubdate_timestamp")
			title = dataJson.GetString("ui_sets.caption_subtitle")
			coverImageLink = dataJson.GetString("cover_landscape")
			content = dataJson.GetString("content")
			rssItem := dao.RSSItem{
				Title:       title,
				Link:        link,
				Description: feed.GenerateDescription(coverImageLink, content),
				Created:     time,
				Thumbnail:   coverImageLink,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "IDAILY_INDEX", rssStr)
	g.Redis().DoVar("EXPIRE", "IDAILY_INDEX", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
