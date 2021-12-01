package _36kr

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"regexp"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"
)

func (ctl *controller) Get36krUserNews(req *ghttp.Request) {
	userId := req.GetString("id")

	if value, err := g.Redis().DoVar("GET", "36KR_USER_"+userId); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://36kr.com/user/" + userId
	rssData := dao.RSSFeed{
		Link:     apiUrl,
		ImageUrl: "https://static.36krcdn.com/36kr-web/static/ic_default_100_56@2x.ec858a2a.png",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		reg := regexp.MustCompile(`"authorDetailData":(.*?),"channel":`)
		contentStr := reg.FindStringSubmatch(resp.ReadAllString())
		if len(contentStr) >= 1 {
			contentData := gjson.New(contentStr[1])
			itemListJson := contentData.GetJsons("authorFlowList.data.itemList")
			authorName := contentData.GetString("authorInfo.data.userNick")
			rssData.Author = authorName
			rssData.Title = "36氪用户 - " + authorName
			rssItems := make([]dao.RSSItem, 0)
			for _, itemJson := range itemListJson {
				widgetContent := itemJson.GetString("templateMaterial.widgetContent")
				widgetImage := itemJson.GetString("templateMaterial.widgetImage")
				description := feed.GenerateDescription(widgetImage, widgetContent)
				rssItem := dao.RSSItem{
					Title:       itemJson.GetString("templateMaterial.widgetTitle"),
					Created:     itemJson.GetString("templateMaterial.publishTime"),
					Description: description,
				}
				router := itemJson.GetString("route")
				if strings.Split(router, "?")[0] == "detail_video" {
					rssItem.Link = "https://36kr.com/video/" + itemJson.GetString("itemId")
				} else {
					rssItem.Link = "https://36kr.com/p/" + itemJson.GetString("itemId")
				}
				rssItems = append(rssItems, rssItem)
			}
			rssData.Items = rssItems
		}

	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "36KR_USER_"+userId, rssStr)
	g.Redis().DoVar("EXPIRE", "36KR_USER_"+userId, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
