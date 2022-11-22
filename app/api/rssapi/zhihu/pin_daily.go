package zhihu

import (
	"context"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetZhihuPinDaily(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	redisKey := "ZHIHU_PIN_DAILY"
	if value, err := component.GetRedis().Do(ctx, "GET", redisKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	hotListUrl := "https://api.zhihu.com/pins/special/972884951192113152/moments?order_by=newest&reverse_order=0&limit=20"
	headers := getHeaders()
	if resp, err := component.GetHttpClient().SetHeaderMap(headers).Get(ctx, hotListUrl); err == nil {
		defer func(resp *gclient.Response) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}(resp)
		rssData := dao.RSSFeed{}
		rssData.Title = "知乎想法-24小时新闻汇总"
		rssData.Link = "https://www.zhihu.com/pin/special/972884951192113152"
		rssData.Description = "汇集每天的社会大事、行业资讯，让你用最简单的方式获得想法里的新闻"
		rssData.Items = getPinRSSItems(resp.ReadAllString())
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		component.GetRedis().Do(ctx, "SET", redisKey, rssStr)
		component.GetRedis().Do(ctx, "EXPIRE", redisKey, 60*60*6)
		req.Response.WriteXmlExit(rssStr)
	}
}
