package infoq

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) GetRecommend(req *ghttp.Request) {
	if value, err := g.Redis().DoVar("GET", "INFOQ_RECOMMEND"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.infoq.cn/hot_recommend.html"
	rssData := dao.RSSFeed{
		Title:       "InfoQ推荐",
		Link:        apiUrl,
		Description: "InfoQ推荐",
		ImageUrl:    "https://static001.infoq.cn/static/infoq/template/img/logo-fasdkjfasdf.png",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		docs := soup.HTMLParse(resp.ReadAllString())
		itemList := docs.FindAll("div", "class", "item-main")
		rssItems := make([]dao.RSSItem, 0)
		for _, item := range itemList {
			imageLink := item.Find("img").Attrs()["src"]
			title := item.Find("a", "class", "com-article-title").Text()
			link := item.Find("a", "class", "com-article-title").Attrs()["href"]
			author := item.Find("a", "class", "com-author-name").Text()
			summary := item.Find("p", "class", "summary").Text()
			rssItem := dao.RSSItem{
				Title:       title,
				Link:        link,
				Description: lib.GenerateDescription(imageLink, summary),
				Author:      author,
			}
			rssItems = append(rssItems, rssItem)
		}

		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData)
	g.Redis().DoVar("SET", "INFOQ_RECOMMEND", rssStr)
	g.Redis().DoVar("EXPIRE", "INFOQ_RECOMMEND", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
