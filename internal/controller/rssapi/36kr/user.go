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
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) Get36krUserNews(req *ghttp.Request) {
	userId := req.Get("id").String()
	var ctx context.Context = context.Background()
	if value, err := cache.GetCache(ctx, "36KR_USER_"+userId); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://36kr.com/user/" + userId
	rssData := dao.RSSFeed{
		Link:     apiUrl,
		ImageUrl: "https://static.36krcdn.com/36kr-web/static/ic_default_100_56@2x.ec858a2a.png",
	}
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		reg := regexp.MustCompile(`"authorDetailData":(.*?),"channel":`)
		contentStr := reg.FindStringSubmatch(resp)
		if len(contentStr) >= 1 {
			contentData := gjson.New(contentStr[1])
			itemListJson := contentData.GetJsons("authorFlowList.data.itemList")
			authorName := contentData.Get("authorInfo.data.userNick").String()
			rssData.Author = authorName
			rssData.Title = "36氪用户 - " + authorName
			rssItems := make([]dao.RSSItem, 0)
			for _, itemJson := range itemListJson {
				widgetContent := itemJson.Get("templateMaterial.widgetContent").String()
				widgetImage := itemJson.Get("templateMaterial.widgetImage").String()
				content := feed.GenerateContent(widgetContent)
				rssItem := dao.RSSItem{
					Title:     itemJson.Get("templateMaterial.widgetTitle").String(),
					Created:   itemJson.Get("templateMaterial.publishTime").String(),
					Content:   content,
					Thumbnail: widgetImage,
				}
				router := itemJson.Get("route").String()
				if strings.Split(router, "?")[0] == "detail_video" {
					rssItem.Link = "https://36kr.com/video/" + itemJson.Get("itemId").String()
				} else {
					rssItem.Link = "https://36kr.com/p/" + itemJson.Get("itemId").String()
				}
				rssItems = append(rssItems, rssItem)
			}
			rssData.Items = rssItems
		}

	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx, "36KR_USER_"+userId, rssStr)
	req.Response.WriteXmlExit(rssStr)
}
