package sciencenet

import (
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]
	cacheKey := "SCIENCE_NET_" + linkConfig.ChannelId
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
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
	if resp := component.GetContent(apiUrl); resp != ""{

		rssItems := commonParser(resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
