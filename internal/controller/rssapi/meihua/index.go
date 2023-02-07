package meihua

import (
	"context"
	"fmt"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]
	cacheKey := "MEIHUA_INDEX" + linkConfig.ChannelId
	if value, err := service.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := fmt.Sprintf("https://www.meihua.info/%s", linkConfig.ChannelId)
	rssData := dao.RSSFeed{
		Title:       "梅花网 -" + linkConfig.Title,
		Link:        apiUrl,
		Tag:         []string{"媒体"},
		Description: "梅花网平台内全部文章内容,梅花网营销行业文章，为营销人提供新鲜、丰富、专业的营销内容、品牌动向、行业趋势。也可自行发布文章，收藏感兴趣的文章。",
		ImageUrl:    "https://www.meihua.info/static/images/icon/meihua.ico",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != ""{
		rssItems := commonParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	service.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	service.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
