package _36kr

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"regexp"
	"rsshub/app/dao"
	"rsshub/lib"
)

type controller struct {
}

var KR36Controller = &controller{}

type NewsRouteConfig struct {
	Link  string
	Title string
	Tags  []string
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
		var (
			summary string
			author  string
			imgLink string
		)
		summary = parseDetail(rssItem.Link)
		author = informationJson.GetString("templateMaterial.authorName")
		imgLink = informationJson.GetString("templateMaterial.widgetImage")
		rssItem.Description = lib.GenerateDescription(imgLink, summary)
		rssItem.Author = author

		rssItems = append(rssItems, rssItem)
	}

	return rssItems
}

func parseDetail(detailLink string) (detailData string) {
	var (
		resp *ghttp.ClientResponse
		err  error
	)
	if resp, err = g.Client().SetHeaderMap(getHeaders()).Get(detailLink); err == nil {
		var (
			reg             *regexp.Regexp
			contentStrArray []string
			contentStr      string
		)
		reg = regexp.MustCompile(`<script>window\.initialState=(.*?)<\/script>`)
		contentStrArray = reg.FindStringSubmatch(resp.ReadAllString())
		if len(contentStrArray) <= 1 {
			g.Log().Errorf("Parse 36kr news detail failed, detail json data is no match rule, detail link is %s", detailLink)
			return
		}
		contentStr = contentStrArray[1]
		contentData := gjson.New(contentStr)

		detailData = contentData.GetString("articleDetail.articleDetailData.data.widgetContent")

	} else {
		g.Log().Errorf("Request 36kr news detail failed, link  %s \nerror : %s", detailLink, err)
	}

	return
}

func getNewsLinks() map[string]NewsRouteConfig {
	Links := map[string]NewsRouteConfig{
		"latest": {
			Link:  "/information/web_news/latest",
			Tags:  []string{"新闻", "互联网", "数据", "金融", "创业", "投资", "快讯", "生活", "科技", "旅游", "职场"},
			Title: "最新"},
		"recommend": {
			Link:  "/information/web_recommend",
			Tags:  []string{"互联网", "数据", "金融", "创业", "投资", "快讯", "生活", "科技", "旅游", "职场"},
			Title: "推荐"},
		"contact": {
			Link:  "/information/contact",
			Tags:  []string{"创业", "投资", "科技", "创投"},
			Title: "创投",
		},
		"ccs": {
			Link:  "/information/ccs",
			Tags:  []string{"互联网", "数据", "金融", "投资", "股票"},
			Title: "中概股",
		},
		"travel": {
			Link:  "/information/travel",
			Tags:  []string{"汽车", "科技"},
			Title: "汽车",
		},
		"technology": {
			Link:  "/technology",
			Tags:  []string{"互联网", "科技", "IT"},
			Title: "科技",
		},
		"enterpriseservice": {
			Link:  "/information/enterpriseservice",
			Tags:  []string{"互联网", "金融", "创业", "企服"},
			Title: "企服",
		},
		"banking": {
			Link:  "/information/banking",
			Tags:  []string{"互联网", "金融", "创业"},
			Title: "金融",
		},
		"life": {
			Link:  "/information/happy_life",
			Tags:  []string{"生活", "互联网"},
			Title: "生活",
		},
		"innovate": {
			Link:  "/information/innovate",
			Tags:  []string{"创新", "互联网"},
			Title: "创新",
		},
		"estate": {
			Link:  "/information/real_estate",
			Tags:  []string{"房产"},
			Title: "房产",
		},
		"workplace": {
			Link:  "/information/web_zhichang",
			Tags:  []string{"职场"},
			Title: "职场",
		},
		"other": {
			Link:  "/information/other",
			Tags:  []string{"其他"},
			Title: "其他",
		},
	}

	return Links
}
