package _199IT

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

type controller struct {
}

var IT1999Controller = &controller{}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func parseArticle(htmlStr string) []dao.RSSItem {
	doc := soup.HTMLParse(htmlStr)
	articleList := doc.FindAll("article")
	rssItems := make([]dao.RSSItem, 0)
	for _, article := range articleList {
		var (
			title       string
			thumbnail   string
			description string
			time        string
			link        string
			detailData  string
		)
		title = article.Find("h2", "class", "entry-title").Find("a").Attrs()["title"]
		thumbnail = article.Find("img", "class", "attachment-post-thumbnail").Attrs()["src"]
		time = article.Find("time", "class", "entry-date").Attrs()["datetime"]
		link = article.Find("h2", "class", "entry-title").Find("a").Attrs()["href"]
		detailData = parseDetail(link)
		description = lib.GenerateDescription(thumbnail, detailData)
		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Description: description,
			Created:     time,
		}
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
			docs        soup.Root
			articleElem soup.Root
		)
		docs = soup.HTMLParse(resp.ReadAllString())
		articleElem = docs.Find("article")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request 199IT article detail failed, link  %s \nerror : %s", detailLink, err)
	}

	return
}
