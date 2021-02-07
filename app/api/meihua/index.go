package meihua

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
	"strings"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]
	cacheKey := "MEIHUA_INDEX" + linkConfig.ChannelId
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := fmt.Sprintf("https://www.meihua.info/%s", linkConfig.ChannelId)
	rssData := dao.RSSFeed{
		Title:       "梅花网 -" + linkConfig.Title,
		Link:        apiUrl,
		Description: "梅花网平台内全部文章内容,梅花网营销行业文章，为营销人提供新鲜、丰富、专业的营销内容、品牌动向、行业趋势。也可自行发布文章，收藏感兴趣的文章。",
		ImageUrl:    "https://www.meihua.info/static/images/icon/meihua.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		rssItems := commonParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
