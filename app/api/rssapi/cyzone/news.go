package cyzone

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *controller) GetNews(req *ghttp.Request) {

	routeArray := strings.Split(req.Router.Uri, "/")
	linkType := routeArray[len(routeArray)-1]
	linkConfig := getNewsLinks()[linkType]

	cacheKey := "CYZONE_INDEX_" + linkType
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
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
	if resp := component.GetContent(apiUrl); resp != "" {
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
			content = parseNewsDetail(link)
			if timeDoc := dataDocs.Find("span", "class", "time"); timeDoc.Error == nil {
				time = timeDoc.Attrs()["data-time"]
			}
			rssItem := dao.RSSItem{
				Title:       title,
				Link:        link,
				Description: feed.GenerateDescription(imageLink, content),
				Created:     time,
				Thumbnail:   imageLink,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*3)
	_ = req.Response.WriteXmlExit(rssStr)
}

func parseNewsDetail(detailLink string) (detailData string) {
	var (
		resp string
	)
	if resp = component.GetContent(detailLink); resp != "" {
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
		g.Log().Errorf("Request cyzone news article detail failed, link  %s \n", detailLink)
	}

	return
}
