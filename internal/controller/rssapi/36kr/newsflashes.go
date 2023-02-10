package _36kr

import (
	"context"
	"regexp"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) Get36krNewsFlashes(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	if value, err := cache.GetCache(ctx, "36KR_NEWS_FLASHES"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://36kr.com/newsflashes"
	rssData := dao.RSSFeed{
		Title:    "快讯 - 36氪",
		Link:     apiUrl,
		ImageUrl: "https://static.36krcdn.com/36kr-web/static/ic_default_100_56@2x.ec858a2a.png",
	}
	if resp := service.GetContent(ctx, apiUrl); resp != "" {

		reg := regexp.MustCompile(`<script>window\.initialState=(.*?)<\/script>`)
		contentStr := reg.FindStringSubmatch(resp)
		if len(contentStr) >= 1 {
			contentStr := contentStr[1]
			contentData := gjson.New(contentStr)
			itemListJson := contentData.GetJsons("newsflashCatalogData.data.newsflashList.data.itemList")
			rssItems := make([]dao.RSSItem, 0)
			for _, itemJson := range itemListJson {
				rssItem := dao.RSSItem{
					Title:   itemJson.Get("templateMaterial.widgetTitle").String(),
					Created: itemJson.Get("templateMaterial.publishTime").String(),
					Content: itemJson.Get("templateMaterial.widgetContent").String(),
				}
				sourceUrlRoute := itemJson.Get("templateMaterial.sourceUrlRoute").String()
				if sourceUrlRoute == "" {
					sourceUrlRoute = "https://36kr.com/newsflashes/" + itemJson.Get("itemId").String()
				} else {
					sourceUrlArr := strings.Split(sourceUrlRoute, "url=")
					if len(sourceUrlArr) > 0 {
						sourceUrlRoute, _ = gurl.Decode(sourceUrlArr[1])
					}
				}
				rssItem.Link = sourceUrlRoute
				rssItems = append(rssItems, rssItem)
			}
			rssData.Items = rssItems
		}
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx, "36KR_NEWS_FLASHES", rssStr)
	req.Response.WriteXmlExit(rssStr)
}
