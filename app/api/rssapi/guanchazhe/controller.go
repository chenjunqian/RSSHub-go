package guanchazhe

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

type Controller struct {
}

type LinkRouteConfig struct {
	ChannelId string
	Title     string
	LinkType  string
	Tags      []string
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func indexParser(htmlStr string) (items []dao.RSSItem) {
	docs := soup.HTMLParse(htmlStr)
	baseUrl := "https://www.guancha.cn"
	imgList := docs.FindAll("ul", "class", "img-List")
	for _, imgItem := range imgList {
		imgItemList := imgItem.FindAll("li")
		if len(imgItemList) > 10 {
			imgItemList = imgItemList[:10]
		}
		for _, newsItem := range imgItemList {
			var imageLink string
			var title string
			var link string
			var author string
			var content string
			var time string

			title = newsItem.Find("h4", "class", "module-title").Find("a").Text()
			imageLink = newsItem.Find("div", "class", "fastRead-img").Find("img").Attrs()["src"]
			link = baseUrl + newsItem.Find("h4", "class", "module-title").Find("a").Attrs()["href"]

			rssItem := dao.RSSItem{
				Title:       title,
				Link:        link,
				Author:      author,
				Description: lib.GenerateDescription(imageLink, content),
				Created:     time,
			}
			items = append(items, rssItem)
		}
	}

	reviewUl := docs.Find("ul", "class", "Review-item")
	if reviewUl.Error != nil {
		return
	}
	reviewList := reviewUl.FindAll("li")
	if len(reviewList) > 10 {
		reviewList = reviewList[:10]
	}
	for _, newsItem := range reviewList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = newsItem.Find("h4", "class", "module-title").Find("a").Text()
		imageLink = newsItem.Find("div", "class", "author-intro").Find("img").Attrs()["src"]
		author = newsItem.Find("div", "class", "author-intro").Find("img").Attrs()["alt"]
		link = baseUrl + newsItem.Find("h4", "class", "module-title").Find("a").Attrs()["href"]
		content = parseCommonDetail(link)

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: lib.GenerateDescription(imageLink, content),
			Created:     time,
		}
		items = append(items, rssItem)
	}

	return
}

func commonParser(htmlStr string) (items []dao.RSSItem) {
	docs := soup.HTMLParse(htmlStr)
	articleList := docs.Find("ul", "class", "new-left-list").FindAll("li")
	baseUrl := "https://www.guancha.cn"
	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.Find("h4", "class", "module-title").Find("a").Text()
		link = baseUrl + article.Find("h4", "class", "module-title").Find("a").Attrs()["href"]
		for _, aTag := range article.FindAll("a") {
			if aTag.Find("img").Error == nil {
				imageLink = aTag.Find("img").Attrs()["src"]
			}
		}
		content = parseCommonDetail(link)
		time = article.Find("span").Text()
		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: lib.GenerateDescription(imageLink, content),
			Created:     time,
		}
		items = append(items, rssItem)
	}
	return
}

func parseCommonDetail(detailLink string) (detailData string) {
	var (
		resp *ghttp.ClientResponse
		err  error
	)
	if resp, err = g.Client().SetHeaderMap(getHeaders()).Get(detailLink); err == nil {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)
		respString = resp.ReadAllString()
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("div", "class", "all-txt")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request guanchazhe article detail failed, link  %s \nerror : %s", detailLink, err)
	}

	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"internation": {
			ChannelId: "internation?s=dhguoji",
			LinkType:  "index",
			Tags:      []string{"海外"},
			Title:     "国际"},
		"military": {
			ChannelId: "military-affairs?s=dhjunshi",
			LinkType:  "index",
			Tags:      []string{"军事"},
			Title:     "军事"},
		"economy": {
			ChannelId: "economy?s=dhcaijing",
			LinkType:  "common",
			Tags:      []string{"财经"},
			Title:     "财经"},
		"tech": {
			ChannelId: "gongye·keji?s=dhgongye·keji",
			LinkType:  "common",
			Tags:      []string{"科技"},
			Title:     "科技"},
		"auto": {
			ChannelId: "qiche?s=dhqiche",
			LinkType:  "common",
			Tags:      []string{"汽车"},
			Title:     "汽车"},
	}

	return Links
}
