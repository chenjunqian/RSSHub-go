package bishijie

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetFlash(req *ghttp.Request) {

	var ctx context.Context = context.Background()
	if value, err := cache.GetCache(ctx, "BISHIJIE_FLASH"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
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
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
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
	cache.SetCache(ctx, "BISHIJIE_FLASH", rssStr)
	req.Response.WriteXmlExit(rssStr)
}
