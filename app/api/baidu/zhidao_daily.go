package baidu

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/encoding/gcharset"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) GetZhiDaoDaily(req *ghttp.Request) {
	if value, err := g.Redis().DoVar("GET", "BAIDU_ZHIDAO_DAILY"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://zhidao.baidu.com/daily?fr=daohang"
	rssData := dao.RSSFeed{
		Title:       "百度知道日报",
		Link:        apiUrl,
		Description: "百度知道日报精选",
		ImageUrl:    "www.baidu.com/favicon.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		respString, _ := gcharset.Convert("UTF-8", "gbk", resp.ReadAllString())
		docs := soup.HTMLParse(respString)
		itemList := docs.FindAll("li", "class", "clearfix")
		rssItems := make([]dao.RSSItem, 0)
		for _, item := range itemList {
			title := item.Find("img").Attrs()["title"]
			imageLink := item.Find("img").Attrs()["src"]
			contentDiv := item.Find("div", "class", "summer")
			var content string
			var link string
			if contentDiv.Error != nil {
				continue
			}

			content = contentDiv.Find("a").FullText()
			link = contentDiv.Find("a").Attrs()["href"]
			link = "https://zhidao.baidu.com/" + link
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
	g.Redis().DoVar("SET", "BAIDU_ZHIDAO_DAILY", rssStr)
	g.Redis().DoVar("EXPIRE", "BAIDU_ZHIDAO_DAILY", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
