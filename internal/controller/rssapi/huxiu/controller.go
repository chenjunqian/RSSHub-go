package huxiu

import (
	"context"
	"rsshub/internal/service"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/frame/g"
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

func parseCommonDetail(ctx context.Context, detailLink string) (detailData string) {
	var (
		resp string
	)
    if resp = service.GetContent(ctx,detailLink); resp != "" {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)

		respString = resp
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("section", "class", "article-wrap")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf(ctx,"Request houxu article detail failed, link  %s \n", detailLink)
	}

	return
}

func getChannelInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"auto": {
			ChannelId: "21",
			Tags:      []string{"汽车"},
			Title:     "车与出行"},
		"young": {
			ChannelId: "106",
			Tags:      []string{"其他"},
			Title:     "年轻一代"},
		"consumer": {
			ChannelId: "103",
			Tags:      []string{"电商"},
			Title:     "电商"},
		"tech": {
			ChannelId: "105",
			Tags:      []string{"科技"},
			Title:     "科技前沿"},
		"finance": {
			ChannelId: "115",
			Tags:      []string{"财经"},
			Title:     "财经 "},
		"entertainment": {
			ChannelId: "22",
			Tags:      []string{"娱乐"},
			Title:     "娱乐淘金"},
		"medical": {
			ChannelId: "111",
			Tags:      []string{"医疗"},
			Title:     "医疗健康"},
		"culture": {
			ChannelId: "113",
			Tags:      []string{"教育"},
			Title:     "文化教育"},
		"oversea": {
			ChannelId: "114",
			Tags:      []string{"海外"},
			Title:     "出海"},
		"realestate": {
			ChannelId: "102",
			Tags:      []string{"金融"},
			Title:     "金融地产"},
		"enterprise": {
			ChannelId: "110",
			Tags:      []string{"企业服务"},
			Title:     "企业服务"},
		"startup": {
			ChannelId: "2",
			Tags:      []string{"创业"},
			Title:     "创业"},
		"social": {
			ChannelId: "112",
			Tags:      []string{"媒体"},
			Title:     "社交"},
		"global": {
			ChannelId: "107",
			Tags:      []string{"海外"},
			Title:     "全球热点"},
		"life": {
			ChannelId: "4",
			Tags:      []string{"生活"},
			Title:     "生活腔调"},
	}

	return Links
}
