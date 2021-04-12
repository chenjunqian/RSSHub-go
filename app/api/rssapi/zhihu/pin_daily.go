package zhihu

import (
	"rsshub/app/dao"
	"rsshub/lib"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetZhihuPinDaily(req *ghttp.Request) {
	redisKey := "ZHIHU_PIN_DAILY"
	if value, err := g.Redis().DoVar("GET", redisKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	hotListUrl := "https://api.zhihu.com/pins/special/972884951192113152/moments?order_by=newest&reverse_order=0&limit=20"
	headers := getHeaders()
	if resp, err := g.Client().SetHeaderMap(headers).Get(hotListUrl); err == nil {
		rssData := dao.RSSFeed{}
		rssData.Title = "知乎想法-24小时新闻汇总"
		rssData.Link = "https://www.zhihu.com/pin/special/972884951192113152"
		rssData.Description = "汇集每天的社会大事、行业资讯，让你用最简单的方式获得想法里的新闻"
		rssData.Items = getPinRSSItems(resp.ReadAllString())
		rssData.ImageUrl = "https://pic4.zhimg.com/80/v2-88158afcff1e7f4b8b00a1ba81171b61_720w.png"
		rssStr := lib.GenerateRSS(rssData)
		g.Redis().DoVar("SET", redisKey, rssStr)
		g.Redis().DoVar("EXPIRE", redisKey, 60*60*6)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
