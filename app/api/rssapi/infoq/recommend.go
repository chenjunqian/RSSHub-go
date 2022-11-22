package infoq

import (
	"context"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetRecommend(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	if value, err := component.GetRedis().Do(ctx,"GET", "INFOQ_RECOMMEND"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.infoq.cn/hot_recommend.html"
	rssData := dao.RSSFeed{
		Title:       "InfoQ推荐",
		Link:        apiUrl,
		Description: "InfoQ推荐",
		Tag:         []string{"汽车"},
		ImageUrl:    "https://static001.infoq.cn/static/infoq/template/img/logo-fasdkjfasdf.png",
	}
	if resp := component.GetContent(ctx,apiUrl); resp != "" {

		docs := soup.HTMLParse(resp)
		itemList := docs.FindAll("div", "class", "item-main")
		rssItems := make([]dao.RSSItem, 0)
		for _, item := range itemList {
			title := item.Find("a", "class", "com-article-title").Text()
			link := item.Find("a", "class", "com-article-title").Attrs()["href"]
			author := item.Find("a", "class", "com-author-name").Text()
			summary := parseRecommendDetail(ctx, link)
			rssItem := dao.RSSItem{
				Title:   title,
				Link:    link,
				Content: feed.GenerateContent(summary),
				Author:  author,
			}
			rssItems = append(rssItems, rssItem)
		}

		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx,"SET", "INFOQ_RECOMMEND", rssStr)
	component.GetRedis().Do(ctx,"EXPIRE", "INFOQ_RECOMMEND", 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
