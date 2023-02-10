package bilibili

import (
	"context"
	"fmt"

	"rsshub/internal/service"
	"rsshub/internal/service/feed"
	"strings"
	"time"

	"rsshub/internal/dao"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetAppVersion(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	id := req.Get("id").String()

	config := map[string]string{
		"android":        "安卓版",
		"iphone":         "iPhone 版",
		"ipad":           "iPad HD 版",
		"win":            "UWP 版",
		"android_tv_yst": "TV 版",
	}

	rootUrl := "https://app.bilibili.com"
	apiUrl := fmt.Sprintf("%s/x/v2/version?mobi_app=%s", rootUrl, id)
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		jsonResp := gjson.New(resp)
		respDataList := jsonResp.GetJsons("data")

		rssData := dao.RSSFeed{}
		rssData.Title = "哔哩哔哩更新情报 - " + config[id]
		rssData.Link = rootUrl

		items := make([]dao.RSSItem, 0)
		for _, dataJson := range respDataList {
			rssItem := dao.RSSItem{
				Link:  rootUrl,
				Title: dataJson.Get("version").String(),
			}
			timeStamp := jsonResp.Get("ptime").Int64()
			rssItem.Created = time.Unix(timeStamp, 0).String()
			desc := dataJson.Get("desc").String()
			descs := strings.Split(desc, "\n-")
			desc = strings.Join(descs, "")
			rssItem.Description = desc
			items = append(items, rssItem)
		}

		rssData.Items = items
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		req.Response.WriteXmlExit(rssStr)
	}
}
