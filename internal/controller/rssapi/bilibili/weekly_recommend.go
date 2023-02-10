package bilibili

import (
	"context"
	"fmt"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"
	"strconv"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetWeeklyRecommend(req *ghttp.Request) {

	var ctx context.Context = context.Background()
	if value, err := cache.GetCache(ctx, "BILIBILI_WEEKLY"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
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
	if statusResp := service.GetContent(ctx, apiUrl); statusResp != "" {
		statusDataJson := gjson.New(statusResp)
		weeklyNumber := statusDataJson.Get("data.0.number").String()
		weeklyName := statusDataJson.Get("data.0.name").String()

		latestApiUrl := "https://app.bilibili.com/x/v2/show/popular/selected?type=weekly_selected&number=" + weeklyNumber
		header["Referer"] = fmt.Sprintf("https://www.bilibili.com/h5/weekly-recommend?num=%s&navhide=1", weeklyName)
		if latestResp := service.GetContent(ctx, latestApiUrl); latestResp != "" {
			latestRespJson := gjson.New(latestResp)
			dataJsonList := latestRespJson.GetJsons("data.list")

			rssItems := make([]dao.RSSItem, 0)
			for _, dataJson := range dataJsonList {
				rssItem := dao.RSSItem{}
				itemTitle := dataJson.Get("title").String()
				rssItem.Title = itemTitle

				recommendReason := dataJson.Get("rcmd_reason")
				cover := dataJson.Get("cover")
				itemDescription := fmt.Sprintf("<br><img src=\"%s\"><br>%s %s<br>%s", cover, weeklyName, itemTitle, recommendReason)
				rssItem.Description = itemDescription

				weeklyNumberFloat, _ := strconv.ParseFloat(weeklyNumber, 64)
				if weeklyNumberFloat > 60 && weeklyNumberFloat > dataJson.Get("bvid").Float64() {
					rssItem.Link = "https://www.bilibili.com/video/" + dataJson.Get("bvid").String()
				} else {
					rssItem.Link = "https://www.bilibili.com/video/av" + dataJson.Get("param").String()
				}
				rssItems = append(rssItems, rssItem)
			}
			rssData.Items = rssItems
		}

	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx, "BILIBILI_WEEKLY", rssStr)
	req.Response.WriteXmlExit(rssStr)
}
