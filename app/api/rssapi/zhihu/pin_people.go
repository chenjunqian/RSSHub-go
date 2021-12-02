package zhihu

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
)

func (ctl *Controller) GetZhihuPinPeople(req *ghttp.Request) {

	peopleId := req.Get("id")

	url := fmt.Sprintf("https://api.zhihu.com/pins/%s/moments?limit=10&offset=0", peopleId)
	headers := getHeaders()
	if resp, err := component.GetHttpClient().SetHeaderMap(headers).Get(url); err == nil {
		respString := resp.ReadAllString()
		jsonResp := gjson.New(respString)
		peopleName := jsonResp.GetString("data.0.target.author.name")
		rssData := dao.RSSFeed{}
		rssData.Title = fmt.Sprintf("%s的知乎想法", peopleName)
		rssData.Link = fmt.Sprintf("https://www.zhihu.com/people/%s/pins", peopleId)
		rssData.Items = getPinRSSItems(respString)
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
