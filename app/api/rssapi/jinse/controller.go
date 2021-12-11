package jinse

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
)

type Controller struct {
}

type LinkRouteConfig struct {
	ChannelId string
	Title     string
	Tags      []string
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func catalogueParser(jsonString string) (items []dao.RSSItem) {
	respJson := gjson.New(jsonString)
	articleList := respJson.GetJsons("list")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.GetString("title")
		imageLink = article.GetString("extra.thumbnail_pic")
		author = article.GetString("extra.author")
		time = article.GetString("extra.published_at")
		link = article.GetString("extra.topic_url")
		content = parseCatalogueDetail(link)

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: feed.GenerateDescription(imageLink, content),
			Created:     time,
			Thumbnail:   imageLink,
		}
		items = append(items, rssItem)
	}
	return
}

func parseCatalogueDetail(detailLink string) (detailData string) {
	var (
		resp *ghttp.ClientResponse
		err  error
	)
	if resp, err = component.GetHttpClient().SetHeaderMap(getHeaders()).Get(detailLink); err == nil {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)
		defer func(resp *ghttp.ClientResponse) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(err)
			}
		}(resp)
		respString = resp.ReadAllString()
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("div", "class", "js-article")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request jinse Catalogue article detail failed, link  %s \nerror : %s", detailLink, err)
	}

	return
}

func livesParser(jsonString string) (items []dao.RSSItem) {
	respJson := gjson.New(jsonString)
	articleList := respJson.GetJsons("list.0.lives")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.GetString("content")
		imageLink = article.GetString("images.0.url")
		time = article.GetString("created_at")
		link = article.GetString("link")
		content = parseLiveDetail(link)

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: feed.GenerateDescription(imageLink, content),
			Created:     time,
			Thumbnail:   imageLink,
		}
		items = append(items, rssItem)
	}
	return
}

func parseLiveDetail(detailLink string) (detailData string) {
	var (
		resp *ghttp.ClientResponse
		err  error
	)
	if resp, err = component.GetHttpClient().SetHeaderMap(getHeaders()).Get(detailLink); err == nil {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)
		defer func(resp *ghttp.ClientResponse) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(err)
			}
		}(resp)
		respString = resp.ReadAllString()
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("section", "class", "at-body")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request jinse Catalogue article detail failed, link  %s \nerror : %s", detailLink, err)
	}

	return
}

func timelineParser(jsonString string) (items []dao.RSSItem) {
	respJson := gjson.New(jsonString)
	articleList := respJson.GetJsons("data.list")

	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.GetString("title")
		imageLink = article.GetString("extra.thumbnail_pic")
		time = article.GetString("extra.published_at")
		link = article.GetString("extra.topic_url")
		author = article.GetString("extra.author")
		content = parseTimelineDetail(link)

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: feed.GenerateDescription(imageLink, content),
			Created:     time,
			Thumbnail:   imageLink,
		}
		items = append(items, rssItem)
	}
	return
}

func parseTimelineDetail(detailLink string) (detailData string) {
	var (
		resp *ghttp.ClientResponse
		err  error
	)
	if resp, err = component.GetHttpClient().SetHeaderMap(getHeaders()).Get(detailLink); err == nil {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)
		defer func(resp *ghttp.ClientResponse) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(err)
			}
		}(resp)
		respString = resp.ReadAllString()
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("div", "class", "js-article")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request jinse Catalogue article detail failed, link  %s \nerror : %s", detailLink, err)
	}

	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"zhengce": {
			ChannelId: "zhengce",
			Tags:      []string{"政治"},
			Title:     "政策"},
		"fenxishishuo": {
			ChannelId: "fenxishishuo",
			Tags:      []string{"财经"},
			Title:     "行情"},
		"defi": {
			ChannelId: "defi",
			Tags:      []string{"财经"},
			Title:     "DeFi"},
		"kuang": {
			ChannelId: "kuang",
			Tags:      []string{"财经"},
			Title:     "矿业"},
		"industry": {
			ChannelId: "industry",
			Tags:      []string{"财经"},
			Title:     "产业"},
		"IPFS": {
			ChannelId: "IPFS",
			Tags:      []string{"财经"},
			Title:     "IPFS"},
		"tech": {
			ChannelId: "tech",
			Tags:      []string{"技术"},
			Title:     "技术"},
		"baike": {
			ChannelId: "baike",
			Tags:      []string{"财经"},
			Title:     "百科"},
		"capitalmarket": {
			ChannelId: "capitalmarket",
			Tags:      []string{"财经"},
			Title:     "研报"},
	}

	return Links
}
