package bishijie

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
	"strings"
)

func (ctl *Controller) GetShenDu(req *ghttp.Request) {

	if value, err := g.Redis().DoVar("GET", "BISHIJIE_SHENDU"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.bishijie.com/shendu"
	rssData := dao.RSSFeed{
		Title:       "币世界-深度",
		Link:        "https://www.baijingapp.com",
		Tag:         []string{"比特币", "金融", "科技", "投资"},
		Description: "币世界网-比特币等数字货币交易所导航、投资理财、快讯、深度、币圈、市场行情第一站。",
		ImageUrl:    "https://www.bishijie.com/favicon.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		docs := soup.HTMLParse(resp.ReadAllString())
		articleDocsList := docs.FindAll("div", "class", "articles-card")
		rssItems := make([]dao.RSSItem, 0)
		for _, articleDoc := range articleDocsList {
			var link string
			var imageLink string
			title := articleDoc.Find("h4", "class", "articles-title").Text()
			title = strings.ReplaceAll(title, "\n", "")
			title = strings.ReplaceAll(title, " ", "")
			articleImgDoc := articleDoc.Find("div", "class", "articles-img")
			if articleImgDoc.Error == nil {
				link = "https://www.bishijie.com" + articleImgDoc.Find("a").Attrs()["href"]
				imageLink = articleImgDoc.Find("img").Attrs()["src"]
			}
			content := articleDoc.Find("p", "class", "articles-text").Text()
			rssItem := dao.RSSItem{
				Title:       title,
				Link:        link,
				Description: lib.GenerateDescription(imageLink, content),
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", "BISHIJIE_SHENDU", rssStr)
	g.Redis().DoVar("EXPIRE", "BISHIJIE_SHENDU", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
