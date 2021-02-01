package idaily

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
	"strconv"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {

	if value, err := g.Redis().DoVar("GET", "IDAILY_INDEX"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "http://idaily-cdn.idailycdn.com/api/list/v3/iphone/zh-hans?page=1&ver=iphone"
	rssData := dao.RSSFeed{
		Title:       "iDaily · 每日环球视野",
		Link:        apiUrl,
		Description: "iDaily 每日环球视野",
		ImageUrl:    "https://dayoneapp.com/favicon-32x32.png?v=9277df7ae7503b6e383587ae0e7210ee",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		respJson := gjson.New(resp.ReadAllString())
		dataJsonArray := respJson.Array()
		rssItems := make([]dao.RSSItem, 0)
		for index := range dataJsonArray {
			dataJson := respJson.GetJson(strconv.Itoa(index))
			link := dataJson.GetString("link_share")
			time := dataJson.GetString("pubdate_timestamp")
			content := dataJson.GetString("content")
			title := dataJson.GetString("ui_sets.caption_subtitle")
			coverImageLink := dataJson.GetString("cover_landscape")
			rssItem := dao.RSSItem{
				Title:       title,
				Link:        link,
				Description: lib.GenerateDescription(coverImageLink, content),
				Created:     time,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", "IDAILY_INDEX", rssStr)
	g.Redis().DoVar("EXPIRE", "IDAILY_INDEX", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
