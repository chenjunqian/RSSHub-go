package cgtn

import (
	"context"
	"fmt"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetMostRead(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]

	if value, err := cache.GetCache(ctx, "CGTN_MOST_"+linkType); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://www.cgtn.com/most-read"
	rssData := dao.RSSFeed{
		Title:    "CGTN - Most " + linkType,
		Link:     apiUrl,
		Tag:      []string{"英文", "海外"},
		ImageUrl: "https://ui.cgtn.com/static/ng/resource/images/logo_title.png",
	}
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		docs := soup.HTMLParse(resp)
		itemWrapper := docs.Find("div", "id", linkType+"Items")
		items := itemWrapper.FindAll("div", "class", "most-read-item-box")
		rssItems := make([]dao.RSSItem, 0)
		for _, item := range items {
			var (
				title       string
				time        string
				link        string
				cgContent   string
				mainContent string
				content     string
			)
			var rssItem = dao.RSSItem{}
			title = item.Find("a").Text()
			time = item.Find("a").Attrs()["data-time"]
			link = item.Find("a").Attrs()["href"]
			cgContent = item.Find("div", "class", "cg-content").Text()
			mainContent = getMainContent(ctx, link)
			content = fmt.Sprintf("%s<br>%s", cgContent, mainContent)
			content = feed.GenerateContent(content)
			rssItem.Title = title
			rssItem.Link = link
			rssItem.Created = time
			rssItem.Content = content
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx, "36KR_NEWS_FLASHES", rssStr)
	req.Response.WriteXmlExit(rssStr)
}
