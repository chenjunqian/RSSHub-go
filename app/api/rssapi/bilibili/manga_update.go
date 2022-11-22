package bilibili

import (
	"context"
	"fmt"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetMangaUpdate(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	id := req.Get("id").String()
	if strings.HasPrefix(id, "mc") {
		id = strings.Replace(id, "mc", "", -1)
	}

	link := "https://manga.bilibili.com/detail/mc" + id
	apiUrl := "https://manga.bilibili.com/twirp/comic.v2.Comic/ComicDetail?device=pc&platform=web"
	header := getHeaders()
	header["Referer"] = link

	rssData := dao.RSSFeed{}
	if resp, err := component.GetHttpClient().SetHeaderMap(header).Post(ctx, apiUrl, g.Map{"comic_id": id}); err == nil {
		respData := gjson.New(resp.ReadAllString())
		dataJson := respData.GetJson("data")
		authorName := dataJson.Get("author_name").Strings()
		author := strings.Join(authorName, ",")
		title := dataJson.Get("title").String()
		rssData.Title = title + " - 哔哩哔哩漫画"
		rssData.Author = author
		rssData.Description = dataJson.Get("classic_lines").String()
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

		epJsonList := dataJson.GetJsons("ep_list")[:20]
		rssItems := make([]dao.RSSItem, 0)
		for _, epJson := range epJsonList {
			rssItem := dao.RSSItem{}
			shortTitle := epJson.Get("short_title").String()
			title := epJson.Get("title").String()
			if shortTitle == title {
				title = shortTitle
			} else {
				title = fmt.Sprintf("%s %s", shortTitle, title)
			}

			description := fmt.Sprintf("<img src='%s'>", epJson.Get("cover"))
			pubDate := epJson.Get("pub_time").String()

			rssItem.Title = title
			rssItem.Author = author
			rssItem.Created = pubDate
			rssItem.Description = description
			rssItem.Link = fmt.Sprintf("https://manga.bilibili.com/mc%s/%s", id, epJson.Get("id"))
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	req.Response.WriteXmlExit(rssStr)
}
