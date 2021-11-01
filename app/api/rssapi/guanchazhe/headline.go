package guanchazhe

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) GetHeadLine(req *ghttp.Request) {

	cacheKey := "GUANCHAZHE_HEADLINE"
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://www.guancha.cn/GuanChaZheTouTiao/list_1.shtml"
	rssData := dao.RSSFeed{
		Title:       "观察者网 - 头条",
		Link:        apiUrl,
		Tag:         []string{"文化", "时事", "政治", "经济", "历史"},
		Description: "观察者网，致力于荟萃中外思想者精华，鼓励青年学人探索，建中西文化交流平台，为崛起中的精英提供决策参考。",
		ImageUrl:    "https://i.guancha.cn/images/favorite.ico",
	}
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		docs := soup.HTMLParse(resp.ReadAllString())
		articleDocList := docs.Find("ul", "class", "headline-list").FindAll("li")
		rssItemList := make([]dao.RSSItem, 0)
		for _, articleDoc := range articleDocList {
			var imageLink string
			var title string
			var link string
			var author string
			var content string
			var time string

			title = articleDoc.Find("h3").Find("a").Text()
			link = "https://www.guancha.cn" + articleDoc.Find("h3").Find("a").Attrs()["href"]
			for _, aTag := range articleDoc.FindAll("a") {
				if aTag.Find("img").Error == nil {
					imageLink = aTag.Find("img").Attrs()["src"]
				}
			}
			content = articleDoc.Find("p", "class", "module-artile").Text()
			time = articleDoc.Find("span").Text()

			rssItem := dao.RSSItem{
				Title:       title,
				Link:        link,
				Author:      author,
				Description: lib.GenerateDescription(imageLink, content),
				Created:     time,
			}
			rssItemList = append(rssItemList, rssItem)
		}
		rssData.Items = rssItemList
	}

	rssStr := lib.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
