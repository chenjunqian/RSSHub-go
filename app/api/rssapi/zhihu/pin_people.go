package zhihu

import (
	"context"
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetZhihuPinPeople(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	peopleId := req.Get("id")

	url := fmt.Sprintf("https://api.zhihu.com/pins/%s/moments?limit=10&offset=0", peopleId)
	headers := getHeaders()
	if resp, err := component.GetHttpClient().SetHeaderMap(headers).Get(ctx,url); err == nil {
		defer func(resp *gclient.Response) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}(resp)
		respString := resp.ReadAllString()
		jsonResp := gjson.New(respString)
		peopleName := jsonResp.Get("data.0.target.author.name")
		rssData := dao.RSSFeed{}
		rssData.Title = fmt.Sprintf("%s的知乎想法", peopleName)
		rssData.Link = fmt.Sprintf("https://www.zhihu.com/people/%s/pins", peopleId)
		rssData.Items = getPinRSSItems(respString)
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		req.Response.WriteXmlExit(rssStr)
	}
}
