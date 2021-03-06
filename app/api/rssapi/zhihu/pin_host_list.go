package zhihu

import (
	"rsshub/app/dao"
	"rsshub/lib"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetZhihuPinHotList(req *ghttp.Request) {
	redisKey := "ZHIHU_PIN_HOT_LIST"
	if value, err := g.Redis().DoVar("GET", redisKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	hotListUrl := "https://api.zhihu.com/pins/hot_list?reverse_order=0"
	headers := getHeaders()
	if resp, err := g.Client().SetHeaderMap(headers).Get(hotListUrl); err == nil {
		rssData := dao.RSSFeed{}
		rssData.Title = "知乎想法热榜"
		rssData.Link = "https://www.zhihu.com/"
		rssData.Description = "每小时更新一次"
		rssData.Items = getPinRSSItems(resp.ReadAllString())
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"
		rssStr := lib.GenerateRSS(rssData)
		g.Redis().DoVar("SET", redisKey, rssStr)
		g.Redis().DoVar("EXPIRE", redisKey, 60*60*6)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
