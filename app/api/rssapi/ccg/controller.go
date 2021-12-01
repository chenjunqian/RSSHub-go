package ccg

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
)

type controller struct {
}

var Controller = &controller{}

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

func indexParser(respString string) (items []dao.RSSItem) {
	respSoup := soup.HTMLParse(respString)
	var articleList []soup.Root
	if articleUl := respSoup.Find("ul", "class", "huodong-list"); articleUl.Error != nil {
		return
	} else {
		articleList = articleUl.FindAll("li")
	}
	for _, article := range articleList {
		var (
			imageLink string
			title     string
			link      string
			author    string
			content   string
			time      string
		)

		if titleH5 := article.Find("h5"); titleH5.Error == nil {
			title = titleH5.Text()
		} else {
			continue
		}

		if imageHtml := article.Find("img"); imageHtml.Error == nil {
			imageLink = imageHtml.Attrs()["src"]
		}

		if linkHtml := article.Find("a"); linkHtml.Error == nil {
			link = linkHtml.Attrs()["href"]
		}

		content = parseIndexDetail(link)

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

func parseIndexDetail(detailLink string) (detailData string) {
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
		articleElem = docs.Find("div", "class", "pinpai-page")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request 199IT article detail failed, link  %s \nerror : %s", detailLink, err)
	}

	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"news": {
			ChannelId: "news",
			Tags:      []string{"新闻", "时事政治"},
			Title:     "新闻动态"},
		"media": {
			ChannelId: "mtbd",
			Tags:      []string{"新闻", "媒体", "时事政治"},
			Title:     "媒体报道"},
	}

	return Links
}
