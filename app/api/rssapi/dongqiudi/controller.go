package dongqiudi

import (
	"rsshub/app/dao"

	"github.com/anaskhan96/soup"
)

type Controller struct {
}

type NewsRouteConfig struct {
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

func commonParser(htmlStr string) (items []dao.RSSItem) {
	respDocs := soup.HTMLParse(htmlStr)
	dataDocsList := respDocs.FindAll("ul")
	if len(dataDocsList) > 15 {
		dataDocsList = dataDocsList[:15]
	}
	for _, dataDocs := range dataDocsList {
		title := dataDocs.Find("a").Text()
		link := "https://www.dongqiudi.com" + dataDocs.Find("a").Attrs()["href"]
		rssItem := dao.RSSItem{
			Title: title,
			Link:  link,
		}
		items = append(items, rssItem)
	}
	return
}

func getNewsLinks() map[string]NewsRouteConfig {
	Links := map[string]NewsRouteConfig{
		"toutiao": {
			ChannelId: "1",
			Tags:      []string{"体育"},
			Title:     "头条"},
		"shendu": {
			ChannelId: "55",
			Tags:      []string{"体育"},
			Title:     "深度"},
		"xianqing": {
			ChannelId: "37",
			Tags:      []string{"体育"},
			Title:     "闲情"},
		"dzhan": {
			ChannelId: "219",
			Tags:      []string{"体育"},
			Title:     "D 站"},
		"zhongchao": {
			ChannelId: "56",
			Tags:      []string{"体育"},
			Title:     "中超"},
		"guoji": {
			ChannelId: "120",
			Tags:      []string{"体育", "海外"},
			Title:     "国际"},
		"yingchao": {
			ChannelId: "3",
			Tags:      []string{"体育", "海外"},
			Title:     "英超"},
		"xijia": {
			ChannelId: "8",
			Tags:      []string{"体育", "海外"},
			Title:     "西甲"},
		"yijia": {
			ChannelId: "4",
			Tags:      []string{"体育", "海外"},
			Title:     "意甲"},
		"dejia": {
			ChannelId: "6",
			Tags:      []string{"体育", "海外"},
			Title:     "德甲"},
		"xinwendabaozha": {
			ChannelId: "41",
			Tags:      []string{"体育"},
			Title:     "新闻大爆炸"},
		"shijiaqiu": {
			ChannelId: "52",
			Tags:      []string{"体育"},
			Title:     "懂球帝十佳球"},
		"mvp": {
			ChannelId: "53",
			Tags:      []string{"体育"},
			Title:     "懂球帝本周 MVP"},
	}

	return Links
}
