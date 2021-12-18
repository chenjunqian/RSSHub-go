package bilibili

import (
	"fmt"
	"rsshub/app/component"
	"rsshub/app/service/feed"
	"strings"
	"time"

	"rsshub/app/dao"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetAppVersion(req *ghttp.Request) {
	id := req.GetString("id")

	config := map[string]string{
		"android":        "安卓版",
		"iphone":         "iPhone 版",
		"ipad":           "iPad HD 版",
		"win":            "UWP 版",
		"android_tv_yst": "TV 版",
	}

	rootUrl := "https://app.bilibili.com"
	apiUrl := fmt.Sprintf("%s/x/v2/version?mobi_app=%s", rootUrl, id)
	if resp := component.GetContent(apiUrl); resp != "" {
		jsonResp := gjson.New(resp)
		respDataList := jsonResp.GetJsons("data")

		rssData := dao.RSSFeed{}
		rssData.Title = "哔哩哔哩更新情报 - " + config[id]
		rssData.Link = rootUrl

		items := make([]dao.RSSItem, 0)
		for _, dataJson := range respDataList {
			rssItem := dao.RSSItem{
				Link:  rootUrl,
				Title: dataJson.GetString("version"),
			}
			timeStamp := jsonResp.GetInt64("ptime")
			rssItem.Created = time.Unix(timeStamp, 0).String()
			desc := dataJson.GetString("desc")
			descs := strings.Split(desc, "\n-")
			desc = strings.Join(descs, "")
			rssItem.Description = desc
			items = append(items, rssItem)
		}

		rssData.Items = items
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
