package baijing

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
)

func (ctl *controller) GetWeekly(req *ghttp.Request) {

	if value, err := g.Redis().DoVar("GET", "BAIJING_WEEKLY"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.baijingapp.com/index/ajax/getlist/type-1"
	rssData := dao.RSSFeed{
		Title:       "白鲸出海-7x24h",
		Link:        "https://www.baijingapp.com/",
		Description: "白鲸出海快讯",
		Tag:         []string{"新闻"},
		ImageUrl:    "https://www.baijingapp.com/static/css/default/img/favicon.ico",
	}
	if resp, err := component.GetHttpClient().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		defer func(resp *ghttp.ClientResponse) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(err)
			}
		}(resp)
		docs := soup.HTMLParse(resp.ReadAllString())
		articleDocList := docs.FindAll("div", "id", "menuKuaixun")
		rssItems := make([]dao.RSSItem, 0)
		for _, articleDoc := range articleDocList {
			title := articleDoc.Find("h3").Text()
			time := articleDoc.Find("span").HTML()
			content := articleDoc.Find("div", "class", "newsflashesText").Text()

			rssItem := dao.RSSItem{
				Title:       title,
				Description: feed.GenerateDescription("", content),
				Author:      "",
				Created:     time,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "BAIJING_WEEKLY", rssStr)
	g.Redis().DoVar("EXPIRE", "BAIJING_WEEKLY", 60*60*8)
	_ = req.Response.WriteXmlExit(rssStr)
}
