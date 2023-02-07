package cyzone

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) GetNews(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getNewsLinks()[linkType]

	cacheKey := "CYZONE_INDEX_" + linkType
	if value, err := service.GetRedis().Do(ctx, "GET", cacheKey); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	var apiUrl string
	if linkConfig.ChannelId == "" {
		apiUrl = "https://www.cyzone.cn/content/channel/index"
	} else {
		apiUrl = "https://www.cyzone.cn/content/channel/index?channel_id=" + linkConfig.ChannelId
	}
	rssData := dao.RSSFeed{
		Title:       "创业邦 - " + linkConfig.Title,
		Link:        apiUrl,
		Tag:         linkConfig.Tags,
		Description: "创业,资本,股权融资,风险投资,VC,IPO,PE,私募,私募股权,上市,融资,天使投资,创业故事,创业项目,投资机构,互联网创业,创业平台",
		ImageUrl:    "https://www.cyzone.cn/favicon.ico",
	}
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		respDocs := soup.HTMLParse(resp)
		dataDocsList := respDocs.FindAll("div", "class", "article-item")
		rssItems := make([]dao.RSSItem, 0)
		for _, dataDocs := range dataDocsList {
			var (
				title     string
				imageLink string
				link      string
				content   string
				time      string
			)
			imageLink = dataDocs.Find("a", "class", "pic-a").Find("img").Attrs()["src"]
			title = dataDocs.Find("a", "class", "item-title").Text()
			link = dataDocs.Find("a", "class", "item-title").Attrs()["href"]
			if strings.HasPrefix(link, "//") {
				link = "https:" + link
			}
			content = parseNewsDetail(ctx, link)
			if timeDoc := dataDocs.Find("span", "class", "time"); timeDoc.Error == nil {
				time = timeDoc.Attrs()["data-time"]
			}
			rssItem := dao.RSSItem{
				Title:     title,
				Link:      link,
				Content:   feed.GenerateContent(content),
				Created:   time,
				Thumbnail: imageLink,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	service.GetRedis().Do(ctx, "SET", cacheKey, rssStr)
	service.GetRedis().Do(ctx, "EXPIRE", cacheKey, 60*60*3)
	req.Response.WriteXmlExit(rssStr)
}

func parseNewsDetail(ctx context.Context, detailLink string) (detailData string) {
	var (
		resp string
	)
	if resp = service.GetContent(ctx, detailLink); resp != "" {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)
		respString = resp
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("div", "class", "show-wrap")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf(ctx, "Request cyzone news article detail failed, link  %s \n", detailLink)
	}

	return
}
