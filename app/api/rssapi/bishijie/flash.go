package bishijie

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetFlash(req *ghttp.Request) {
	if value, err := g.Redis().DoVar("GET", "BISHIJIE_FLASH"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.bishijie.com/kuaixun/"
	rssData := dao.RSSFeed{
		Title:       "币世界-快讯",
		Link:        "https://www.bishijie.com",
		Tag:         []string{"比特币", "金融", "科技", "投资", "新闻"},
		Description: "币世界网-比特币等数字货币交易所导航、投资理财、快讯、深度、币圈、市场行情第一站。",
		ImageUrl:    "https://www.bishijie.com/favicon.ico",
	}
	if resp := component.GetContent(apiUrl); resp != "" {
		docs := soup.HTMLParse(resp)
		newsContainer := docs.Find("ul", "class", "newscontainer")
		dataListDocs := newsContainer.FindAll("li")
		rssItems := make([]dao.RSSItem, 0)
		for _, dataDocs := range dataListDocs {
			contentDocs := dataDocs.Find("div", "class", "content")
			var title string
			var link string
			var content string
			var imageLink string
			var description string
			if contentDocs.Error == nil {
				title = contentDocs.Find("h3").Text()
				title = strings.ReplaceAll(title, "\n", "")
				title = strings.ReplaceAll(title, " ", "")
				link = "https://www.bishijie.com" + contentDocs.Find("a").Attrs()["href"]
				newsSubDiv := contentDocs.Find("div", "class", "news-content")
				content = newsSubDiv.Find("div").Text()
				imageDiv := newsSubDiv.Find("img")
				if imageDiv.Error == nil {
					imageLink = imageDiv.Attrs()["src"]
					content = feed.GenerateContent(content)
				}
			}
			rssItem := dao.RSSItem{
				Title:       title,
				Link:        link,
				Description: description,
				Thumbnail:   imageLink,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "BISHIJIE_FLASH", rssStr)
	g.Redis().DoVar("EXPIRE", "BISHIJIE_FLASH", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
