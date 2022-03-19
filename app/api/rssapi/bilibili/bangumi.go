package bilibili

import (
	"fmt"
	"regexp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetBangumi(req *ghttp.Request) {
	mediaId := req.GetString("id")
	apiUrl := "https://www.bilibili.com/bangumi/media/md" + mediaId
	if resp := component.GetContent(apiUrl); resp != "" {
		reg := regexp.MustCompile(`window\.__INITIAL_STATE__=([\s\S]+);\(function\(\)`)
		contentStrs := reg.FindStringSubmatch(resp)
		if len(contentStrs) <= 1 {
			_ = req.Response.WriteXmlExit("")
		}
		contentStr := contentStrs[1]
		contentData := gjson.New(contentStr)
		seasonId := contentData.GetString("mediaInfo.season_id")

		rssData := dao.RSSFeed{}
		items := make([]dao.RSSItem, 0)
		seasonUrl := "https://api.bilibili.com/pgc/web/season/section?season_id=" + seasonId
		if seasonResp := component.GetContent(seasonUrl); resp != "" {
			seasonJsonResp := gjson.New(seasonResp)
			seasonData := seasonJsonResp.GetJson("result")

			if seasonData.Get("main_section.episodes") != nil {
				episodeJsons := seasonData.GetJsons("main_section.episodes")
				for _, episodeJson := range episodeJsons {
					rssItem := dao.RSSItem{}
					rssItem.Title = fmt.Sprintf("第%s话 %s", episodeJson.GetString("title"), episodeJson.GetString("long_title"))
					rssItem.Content = fmt.Sprintf("<img src='%s'>", episodeJson.GetString("cover"))
					rssItem.Link = "https://www.bilibili.com/bangumi/play/ep" + episodeJson.GetString("id")
					items = append(items, rssItem)
				}
			}

			if seasonData.Get("section") != nil {
				sectionJsons := seasonData.GetJsons("section")
				for _, sectionJson := range sectionJsons {
					rssItem := dao.RSSItem{}
					rssItem.Title = fmt.Sprintf("%s %s", sectionJson.GetString("title"), sectionJson.GetString("long_title"))
					rssItem.Content = fmt.Sprintf("<img src='%s'>", sectionJson.GetString("cover"))
					rssItem.Link = "https://www.bilibili.com/bangumi/play/ep" + sectionJson.GetString("id")
					items = append(items, rssItem)
				}
			}
		}

		rssData.Link = "https://www.bilibili.com/bangumi/media/md" + contentData.GetString("mediaInfo.media_id")
		rssData.Title = contentData.GetString("mediaInfo.title")
		rssData.Description = contentData.GetString("mediaInfo.evaluate")
		rssData.Items = items
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		_ = req.Response.WriteXmlExit(rssStr)
	}
}
