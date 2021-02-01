package _36kr

import (
	"github.com/gogf/gf/encoding/gjson"
	"regexp"
	"rsshub/app/dao"
	"rsshub/lib"
)

type Controller struct {
}

type NewsRouteConfig struct {
	Link  string
	Title string
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func parseNews(htmlStr string) []dao.RSSItem {
	rssItems := make([]dao.RSSItem, 0)
	reg := regexp.MustCompile(`<script>window\.initialState=(.*?)<\/script>`)
	contentStrs := reg.FindStringSubmatch(htmlStr)
	if len(contentStrs) <= 1 {
		return rssItems
	}
	contentStr := contentStrs[1]
	contentData := gjson.New(contentStr)
	informationList := contentData.GetJsons("information.informationList.itemList")
	for _, informationJson := range informationList {
		rssItem := dao.RSSItem{
			Title:   informationJson.GetString("templateMaterial.widgetTitle"),
			Link:    "https://36kr.com/p/" + informationJson.GetString("itemId"),
			Created: informationJson.GetString("templateMaterial.publishTime"),
		}
		summary := informationJson.GetString("templateMaterial.summary")
		author := informationJson.GetString("templateMaterial.authorName")
		imgLink := informationJson.GetString("templateMaterial.widgetImage")
		rssItem.Description = lib.GenerateDescription(imgLink, summary)
		rssItem.Author = author

		rssItems = append(rssItems, rssItem)
	}

	return rssItems
}

func getNewsLinks() map[string]NewsRouteConfig {
	Links := map[string]NewsRouteConfig{
		"latest": {
			Link:  "/information/web_news/latest",
			Title: "最新"},
		"recommend": {
			Link:  "/information/web_recommend",
			Title: "推荐"},
		"contact": {
			Link:  "/information/contact",
			Title: "创投",
		},
		"ccs": {
			Link:  "/information/ccs",
			Title: "中概股",
		},
		"travel": {
			Link:  "/information/travel",
			Title: "汽车",
		},
		"technology": {
			Link:  "/technology",
			Title: "科技",
		},
		"enterpriseservice": {
			Link:  "/information/enterpriseservice",
			Title: "企服",
		},
		"banking": {
			Link:  "/information/banking",
			Title: "金融",
		},
		"life": {
			Link:  "/information/happy_life",
			Title: "生活",
		},
		"innovate": {
			Link:  "/information/innovate",
			Title: "创新",
		},
		"estate": {
			Link:  "/information/real_estate",
			Title: "房产",
		},
		"workplace": {
			Link:  "/information/web_zhichang",
			Title: "职场",
		},
		"other": {
			Link:  "/information/other",
			Title: "其他",
		},
	}

	return Links
}
