package sspai

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
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func commonParser(respString string) (items []dao.RSSItem) {
	respJson := gjson.New(respString)
	articleList := respJson.GetJsons("data")
	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		title = article.GetString("title")
		imageLink = "https://cdn.sspai.com/" + article.GetString("banner")
		author = article.GetString("author.nickname")
		time = article.GetString("created_time")
		link = "https://sspai.com/post/" + article.GetString("id")
		content = parseCommonDetail(link)

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

func parseCommonDetail(detailLink string) (detailData string) {
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
		articleElem = docs.Find("article", "class", "normal-article")
		if articleElem.Pointer == nil {
			return respString
		}
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request sspai index article detail failed, link  %s \nerror : %s", detailLink, err)
	}

	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"recommend": {
			ChannelId: "recommend",
			Title:     "推荐"},
		"hot": {
			ChannelId: "hot",
			Title:     "热门文章"},
		"app_recommend": {
			ChannelId: "app_recommend",
			Title:     "应用推荐"},
		"skill": {
			ChannelId: "skill",
			Title:     "效率技巧"},
		"lifestyle": {
			ChannelId: "lifestyle",
			Title:     "生活方式"},
		"podcast": {
			ChannelId: "podcast",
			Title:     "少数派播客"},
	}

	return Links
}
