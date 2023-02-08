package jinse

import (
	"context"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetLives(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	cacheKey := "JINSE_LIVES"
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://api.jinse.com/v4/live/list?limit=20"
	rssData := dao.RSSFeed{
		Title:       "金色财经 - 直播",
		Link:        "https://www.jinse.com/",
		Tag:         []string{"财经"},
		Description: "金色财经是集行业新闻、资讯、行情、数据等一站式区块链产业服务平台，我们追求及时、全面、专业、准确的资讯与数据，致力于为区块链创业者以及数字货币投资者提供最好的产品和服务。",
		ImageUrl:    "https://www.jinse.com/favicon.ico",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != ""{

		rssItems := livesParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx,cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
