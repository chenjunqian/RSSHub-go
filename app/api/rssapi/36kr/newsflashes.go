package _36kr

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gurl"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"regexp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"
)

func (ctl *controller) Get36krNewsFlashes(req *ghttp.Request) {

	if value, err := g.Redis().DoVar("GET", "36KR_NEWS_FLASHES"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://36kr.com/newsflashes"
	rssData := dao.RSSFeed{
		Title:    "快讯 - 36氪",
		Link:     apiUrl,
		ImageUrl: "https://static.36krcdn.com/36kr-web/static/ic_default_100_56@2x.ec858a2a.png",
	}
	if resp, err := component.GetHttpClient().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {

		reg := regexp.MustCompile(`<script>window\.initialState=(.*?)<\/script>`)
		contentStr := reg.FindStringSubmatch(resp.ReadAllString())
		if len(contentStr) >= 1 {
			contentStr := contentStr[1]
			contentData := gjson.New(contentStr)
			itemListJson := contentData.GetJsons("newsflashCatalogData.data.newsflashList.data.itemList")
			rssItems := make([]dao.RSSItem, 0)
			for _, itemJson := range itemListJson {
				rssItem := dao.RSSItem{
					Title:       itemJson.GetString("templateMaterial.widgetTitle"),
					Created:     itemJson.GetString("templateMaterial.publishTime"),
					Description: itemJson.GetString("templateMaterial.widgetContent"),
				}
				sourceUrlRoute := itemJson.GetString("templateMaterial.sourceUrlRoute")
				if sourceUrlRoute == "" {
					sourceUrlRoute = "https://36kr.com/newsflashes/" + itemJson.GetString("itemId")
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
	g.Redis().DoVar("SET", "36KR_NEWS_FLASHES", rssStr)
	g.Redis().DoVar("EXPIRE", "36KR_NEWS_FLASHES", 60*60*1)
	_ = req.Response.WriteXmlExit(rssStr)
}
