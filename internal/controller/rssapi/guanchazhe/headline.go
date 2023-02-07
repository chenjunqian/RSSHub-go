package guanchazhe

import (
	"context"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetHeadLine(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	cacheKey := "GUANCHAZHE_HEADLINE"
	if value, err := service.GetRedis().Do(ctx,"GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
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
	if resp := service.GetContent(ctx,apiUrl); resp != "" {

		docs := soup.HTMLParse(resp)
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
			content = parseCommonDetail(ctx, link)
			time = articleDoc.Find("span").Text()

			rssItem := dao.RSSItem{
				Title:     title,
				Link:      link,
				Author:    author,
				Content:   feed.GenerateContent(content),
				Created:   time,
				Thumbnail: imageLink,
			}
			rssItemList = append(rssItemList, rssItem)
		}
		rssData.Items = rssItemList
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	service.GetRedis().Do(ctx,"SET", cacheKey, rssStr)
	service.GetRedis().Do(ctx,"EXPIRE", cacheKey, 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
