package ccg

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetIndex(req *ghttp.Request) {

	var ctx context.Context = context.Background()
	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]

	cacheKey := "CCG_INDEX_" + linkConfig.ChannelId
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "http://www.ccg.org.cn/" + linkConfig.ChannelId
	rssData := dao.RSSFeed{
		Title:       "全球化智库 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         linkConfig.Tags,
		Description: "以全球的视野为中国建言，以中国智慧为全球贡献。",
		ImageUrl:    "http://www.ccg.org.cn/favicon.ico",
	}

	if resp := service.GetContent(ctx,apiUrl); resp != "" {
		rssItems := indexParser(ctx,resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx,cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
