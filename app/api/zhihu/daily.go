package zhihu

import (
	"rsshub/app/dao"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetDaily(req *ghttp.Request) {

	dailyUrl := "https://news-at.zhihu.com/api/4/news/latest"
	headers := getHeaders()
	headers["Referer"] = dailyUrl
	if resp, err := g.Client().SetHeaderMap(headers).Get(dailyUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		respDataList := jsonResp.GetArray("data")

		rssData := dao.RSSFeed{}
		rssData.Title = "知乎日报"
		rssData.Link = dailyUrl
		rssData.Description = "每天3次，每次7分钟"
		for index := range respDataList {

		}
	}
}
