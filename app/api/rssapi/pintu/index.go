package pintu

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"
)

func (ctl *Controller) GetIndex(req *ghttp.Request) {

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getInfoLinks()[linkType]

	cacheKey := "PINTU_INDEX_" + linkConfig.ChannelId
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.pintu360.com/service/ajax_article_service.php"
	rssData := dao.RSSFeed{
		Title:       "品途 - " + linkConfig.Title,
		Link:        apiUrl,
		Description: "品途商业评论是深扎智能科技,文娱,新零售,新消费,创业投资,大健康等创新产业的评论型媒体,是新经济，新商业，新产业的风向标;致力用媒体激发产业创新,赋能相关产业互联网化转型升级",
		ImageUrl:    "https://www.niaogebiji.com/favicon.ico",
	}
	var contentType string
	if linkConfig.ChannelId == "0" {
		contentType = "recommend"
	} else {
		contentType = "classId"
	}
	if resp, err := component.GetHttpClient().SetHeaderMap(getHeaders()).Post(apiUrl, g.Map{
		"fnName":     "getArticleList",
		"type":       contentType,
		"id":         linkConfig.ChannelId,
		"pageNumber": 0,
		"duration":   "quarter",
	}); err == nil {
		defer func(resp *ghttp.ClientResponse) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(err)
			}
		}(resp)
		rssItems := indexParser(resp.ReadAllString())
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
