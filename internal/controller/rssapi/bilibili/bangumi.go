package bilibili

import (
	"context"
	"fmt"
	"regexp"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetBangumi(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	mediaId := req.Get("id").String()
	apiUrl := "https://www.bilibili.com/bangumi/media/md" + mediaId
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		reg := regexp.MustCompile(`window\.__INITIAL_STATE__=([\s\S]+);\(function\(\)`)
		contentStrs := reg.FindStringSubmatch(resp)
		if len(contentStrs) <= 1 {
			req.Response.WriteXmlExit("")
		}
		contentStr := contentStrs[1]
		contentData := gjson.New(contentStr)
		seasonId := contentData.Get("mediaInfo.season_id").String()

		rssData := dao.RSSFeed{}
		items := make([]dao.RSSItem, 0)
		seasonUrl := "https://api.bilibili.com/pgc/web/season/section?season_id=" + seasonId
		if seasonResp := service.GetContent(ctx, seasonUrl); resp != "" {
			seasonJsonResp := gjson.New(seasonResp)
			seasonData := seasonJsonResp.GetJson("result")

			if seasonData.Get("main_section.episodes") != nil {
				episodeJsons := seasonData.GetJsons("main_section.episodes")
				for _, episodeJson := range episodeJsons {
					rssItem := dao.RSSItem{}
					rssItem.Title = fmt.Sprintf("第%s话 %s", episodeJson.Get("title"), episodeJson.Get("long_title"))
					rssItem.Content = fmt.Sprintf("<img src='%s'>", episodeJson.Get("cover"))
					rssItem.Link = "https://www.bilibili.com/bangumi/play/ep" + episodeJson.Get("id").String()
					items = append(items, rssItem)
				}
			}

			if seasonData.Get("section") != nil {
				sectionJsons := seasonData.GetJsons("section")
				for _, sectionJson := range sectionJsons {
					rssItem := dao.RSSItem{}
					rssItem.Title = fmt.Sprintf("%s %s", sectionJson.Get("title"), sectionJson.Get("long_title"))
					rssItem.Content = fmt.Sprintf("<img src='%s'>", sectionJson.Get("cover"))
					rssItem.Link = "https://www.bilibili.com/bangumi/play/ep" + sectionJson.Get("id").String()
					items = append(items, rssItem)
				}
			}
		}

		rssData.Link = "https://www.bilibili.com/bangumi/media/md" + contentData.Get("mediaInfo.media_id").String()
		rssData.Title = contentData.Get("mediaInfo.title").String()
		rssData.Description = contentData.Get("mediaInfo.evaluate").String()
		rssData.Items = items
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"
		rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
		req.Response.WriteXmlExit(rssStr)
	}
}
