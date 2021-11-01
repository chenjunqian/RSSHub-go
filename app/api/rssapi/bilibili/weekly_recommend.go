package bilibili

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
	"strconv"
)

func (ctl *Controller) GetWeeklyRecommend(req *ghttp.Request) {

	if value, err := g.Redis().DoVar("GET", "BILIBILI_WEEKLY"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://app.bilibili.com/x/v2/show/popular/selected/series?type=weekly_selected"
	header := getHeaders()
	header["Referer"] = "https://www.bilibili.com/h5/weekly-recommend"
	rssData := dao.RSSFeed{}
	rssData.Description = "B站每周必看"
	rssData.Title = "B站每周必看"
	rssData.Link = "https://www.bilibili.com/h5/weekly-recommend"
	rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"
	if statusResp, err := g.Client().SetHeaderMap(header).Get(apiUrl); err == nil {
		statusDataJson := gjson.New(statusResp.ReadAllString())
		weeklyNumber := statusDataJson.GetString("data.0.number")
		weeklyName := statusDataJson.GetString("data.0.name")

		latestApiUrl := "https://app.bilibili.com/x/v2/show/popular/selected?type=weekly_selected&number=" + weeklyNumber
		header["Referer"] = fmt.Sprintf("https://www.bilibili.com/h5/weekly-recommend?num=%s&navhide=1", weeklyName)
		if latestResp, err := g.Client().SetHeaderMap(header).Get(latestApiUrl); err == nil {
			latestRespJson := gjson.New(latestResp.ReadAllString())
			dataJsonList := latestRespJson.GetJsons("data.list")

			rssItems := make([]dao.RSSItem, 0)
			for _, dataJson := range dataJsonList {
				rssItem := dao.RSSItem{}
				itemTitle := dataJson.GetString("title")
				rssItem.Title = itemTitle

				recommendReason := dataJson.GetString("rcmd_reason")
				cover := dataJson.GetString("cover")
				itemDescription := fmt.Sprintf("<br><img src=\"%s\"><br>%s %s<br>%s", cover, weeklyName, itemTitle, recommendReason)
				rssItem.Description = itemDescription

				weeklyNumberFloat, _ := strconv.ParseFloat(weeklyNumber, 64)
				if weeklyNumberFloat > 60 && weeklyNumberFloat > dataJson.GetFloat64("bvid") {
					rssItem.Link = "https://www.bilibili.com/video/" + dataJson.GetString("bvid")
				} else {
					rssItem.Link = "https://www.bilibili.com/video/av" + dataJson.GetString("param")
				}
				rssItems = append(rssItems, rssItem)
			}
			rssData.Items = rssItems
		}

	}
	rssStr := lib.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "BILIBILI_WEEKLY", rssStr)
	g.Redis().DoVar("EXPIRE", "BILIBILI_WEEKLY", 60*60*3)
	_ = req.Response.WriteXmlExit(rssStr)
}
