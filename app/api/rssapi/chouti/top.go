package chouti

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *controller) GetTop(req *ghttp.Request) {
	cacheKey := "CHOUTI_TOP"
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://dig.chouti.com"
	rssData := dao.RSSFeed{
		Title:       "抽屉新热榜Top",
		Link:        "https://dig.chouti.com",
		Tag:         []string{"社区", "门户", "其他"},
		Description: "抽屉新热榜，汇聚每日搞笑段子、热门图片、有趣新闻。它将微博、门户、社区、bbs、社交网站等海量内容聚合在一起，通过用户推荐生成最热榜单。看抽屉新热榜，每日热门、有趣资讯尽收眼底",
		ImageUrl:    "https://m.chouti.com/static/image/favicon.png",
	}
	if resp := component.GetContent(apiUrl); resp != "" {
		respDocs := soup.HTMLParse(resp)
		dataDocsList := respDocs.FindAll("div", "class", "link-area")
		rssItems := make([]dao.RSSItem, 0)
		for _, dataDocs := range dataDocsList {
			var (
				title     string
				link      string
				imageLink string
				time      string
				author    string
				content   string
			)
			title = dataDocs.Find("a", "class", "link-title").Text()
			linkFrom := dataDocs.Find("div", "class", "link-from")

			if linkFrom.Error != nil || linkFrom.Text() == "" {
				link = apiUrl + dataDocs.Find("a", "class", "link-title").Attrs()["href"]
			} else {
				link = dataDocs.Find("a", "class", "link-title").Attrs()["href"]
			}

			if imageLinkDoc := dataDocs.Find("img", "class", "image-scale"); imageLinkDoc.Error == nil {
				imageLink = imageLinkDoc.Attrs()["src"]
			}
			time = dataDocs.Find("span", "class", "time-update").Attrs()["data-time"]
			author = dataDocs.Find("span", "class", "author-name").Text()
			if content == "" {
				content = imageLink
			}
			rssItem := dao.RSSItem{
				Title:     title,
				Link:      link,
				Content:   feed.GenerateContent(content),
				Author:    author,
				Created:   time,
				Thumbnail: imageLink,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*3)
	_ = req.Response.WriteXmlExit(rssStr)
}
