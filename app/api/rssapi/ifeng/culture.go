package ifeng

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"
)

func (ctl *Controller) GetCulture(req *ghttp.Request) {
	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getCultureInfoLinks()[linkType]

	cacheKey := "IFENG_ENTERTAINMENT_" + linkConfig.ChannelId
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://ent.ifeng.com/" + linkConfig.ChannelId + "/"
	rssData := dao.RSSFeed{
		Title:       "凤凰网|文化 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         []string{"文化"},
		Description: "凤凰新媒体不仅是控股的凤凰卫视传媒集团优质电视内容的网络传播渠道，更整合了众多专业媒体机构生产的内容、用户生成的内容、以及自身生产的专业内容，提供含图文音视频的全方位综合新闻资讯、深度报道、观点评论、财经产品、互动应用、分享社区、在线网页游戏等服务，满足主流人群浏览、表达、交流、分享、娱乐、理财等多元化与个性化的诉求，并反向传输给凤凰卫视的电视平台，形成创新的网台联动组合传播模式，为互联网、移动互联网及视频用户提供丰富的内容与随时随地随身的服务。",
		ImageUrl:    "https://y0.ifengimg.com/index/favicon.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssData.Items = commonParser(resp.ReadAllString())
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
