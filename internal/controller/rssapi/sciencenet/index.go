package sciencenet

import (
	"context"
	"fmt"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]
	cacheKey := "SCIENCE_NET_" + linkConfig.ChannelId
	if value, err := cache.GetCache(ctx, cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := fmt.Sprintf("http://blog.sciencenet.cn/blog.php?mod=%s&type=list&op=1&ord=1", linkConfig.ChannelId)
	rssData := dao.RSSFeed{
		Title:       "科学网博客 -" + linkConfig.Title,
		Link:        "http://blog.sciencenet.cn/",
		Tag:         []string{"科技"},
		Description: "科学网博客-构建全球华人科学博客圈",
		ImageUrl:    "https://blog.sciencenet.cn/favicon.ico",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != ""{

		rssItems := commonParser(ctx, resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx,cacheKey, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
