package meihua

import (
	"context"
	"regexp"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/frame/g"
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

func commonParser(ctx context.Context, respString string) (items []dao.RSSItem) {
	respDoc := soup.HTMLParse(respString)
	articleList := respDoc.FindAll("div", "class", "article-item")
	baseUrl := "https://www.meihua.info"
	if len(articleList) > 10 {
		articleList = articleList[:10]
	}
	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		coverDiv := article.Find("div", "class", "article-item-cover")
		imgDoc := coverDiv.Find("img")
		title = imgDoc.Attrs()["alt"]
		coverDivString := coverDiv.HTML()
		reg := regexp.MustCompile(`<noscript><img src="(.*?)"`)
		imageUrlMatches := reg.FindStringSubmatch(coverDivString)
		if len(imageUrlMatches) > 1 {
			imageLink = imageUrlMatches[1]
		}

		link = baseUrl + coverDiv.Find("a").Attrs()["href"]

		content = parseCommonDetail(ctx, link)
		spanList := article.FindAll("span", "class", "text-tag")
		for _, spanDoc := range spanList {
			if aTag := spanDoc.Find("a"); aTag.Error == nil {
				author = aTag.Text()
			}
			if spanTag := spanDoc.Find("span"); spanTag.Error == nil {
				time = spanTag.Text()
			}
		}

		rssItem := dao.RSSItem{
			Title:     title,
			Link:      link,
			Author:    author,
			Content:   feed.GenerateContent(content),
			Created:   time,
			Thumbnail: imageLink,
		}
		items = append(items, rssItem)
	}
	return
}

func parseCommonDetail(ctx context.Context, detailLink string) (detailData string) {
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
		articleElem = docs.Find("section", "id", "article-content-html")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf(ctx, "Request meihua article detail failed, link  %s \n", detailLink)
	}

	return
}

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"latest": {
			ChannelId: "article!2",
			Title:     "最新"},
		"hot": {
			ChannelId: "article!3",
			Title:     "热门"},
	}

	return Links
}
