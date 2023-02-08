package zhihu

import (
	"context"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetZhihuPinHotList(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	redisKey := "ZHIHU_PIN_HOT_LIST"
	if value, err := cache.GetCache(ctx, redisKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	hotListUrl := "https://api.zhihu.com/pins/hot_list?reverse_order=0"
	headers := getHeaders()
	if resp, err := service.GetHttpClient().SetHeaderMap(headers).Get(ctx, hotListUrl); err == nil {
		defer func(resp *gclient.Response) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}(resp)
		rssData := dao.RSSFeed{}
		rssData.Title = "知乎想法热榜"
		rssData.Link = "https://www.zhihu.com/"
		rssData.Description = "每小时更新一次"
		rssData.Items = getPinRSSItems(resp.ReadAllString())
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		cache.SetCache(ctx,redisKey, rssStr)
		req.Response.WriteXmlExit(rssStr)
	}
}
