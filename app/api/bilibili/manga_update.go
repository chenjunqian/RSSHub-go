package bilibili

import (
	"fmt"
	"rsshub/app/dao"
	"rsshub/lib"
	"strings"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetMangaUpdate(req *ghttp.Request) {
	id := req.GetString("id")
	if strings.HasPrefix(id, "mc") {
		id = strings.Replace(id, "mc", "", -1)
	}

	link := "https://manga.bilibili.com/detail/mc" + id
	apiUrl := "https://manga.bilibili.com/twirp/comic.v2.Comic/ComicDetail?device=pc&platform=web"
	header := getHeaders()
	header["Referer"] = link

	rssData := dao.RSSFeed{}
	if resp, err := g.Client().SetHeaderMap(header).Post(apiUrl, g.Map{"comic_id": id}); err == nil {
		respData := gjson.New(resp.ReadAllString())
		dataJson := respData.GetJson("data")
		authorName := dataJson.GetStrings("author_name")
		author := strings.Join(authorName, ",")
		title := dataJson.GetString("title")
		rssData.Title = title + " - 哔哩哔哩漫画"
		rssData.Author = author
		rssData.Description = dataJson.GetString("classic_lines")
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

		epJsonList := dataJson.GetJsons("ep_list")[:20]
		rssItems := make([]dao.RSSItem, 0)
		for _, epJson := range epJsonList {
			rssItem := dao.RSSItem{}
			shortTitle := epJson.GetString("short_title")
			title := epJson.GetString("title")
			if shortTitle == title {
				title = shortTitle
			} else {
				title = fmt.Sprintf("%s %s", shortTitle, title)
			}

			description := fmt.Sprintf("<img src='%s'>", epJson.GetString("cover"))
			pubDate := epJson.GetString("pub_time")

			rssItem.Title = title
			rssItem.Author = author
			rssItem.Created = pubDate
			rssItem.Description = description
			rssItem.Link = fmt.Sprintf("https://manga.bilibili.com/mc%s/%s", id, epJson.GetString("id"))
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData)
	_ = req.Response.WriteXmlExit(rssStr)
}
