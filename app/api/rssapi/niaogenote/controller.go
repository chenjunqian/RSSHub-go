package niaogenote

import (
	"github.com/anaskhan96/soup"
	"regexp"
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

func catParser(respString string) (items []dao.RSSItem) {
	respDoc := soup.HTMLParse(respString)
	articleList := respDoc.FindAll("div", "class", "articleBox")
	baseUrl := "https://www.niaogebiji.com"
	for _, article := range articleList {
		var imageLink string
		var title string
		var link string
		var author string
		var content string
		var time string

		if imageATag := article.Find("a", "class", "articleImgLink"); imageATag.Error == nil {
			styleString := imageATag.Attrs()["style"]
			reg := regexp.MustCompile(`url\('(.*?)'\)`)
			contentStrArr := reg.FindStringSubmatch(styleString)
			if len(contentStrArr) <= 1 {
				return
			}
			imageLink = contentStrArr[1]
			title = imageATag.Text()
			link = baseUrl + imageATag.Attrs()["href"]
		}

		if contentATag := article.Find("a", "class", "articleContentInner"); contentATag.Error == nil {
			content = contentATag.Text()
		}

		if timeTag := article.Find("span", "class", "writeTime"); timeTag.Error == nil {
			time = timeTag.Text()
		}

		if timeTag := article.Find("span", "class", "writeTime"); timeTag.Error == nil {
			time = timeTag.Text()
		}

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

func getInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"user_op": {
			ChannelId: "101",
			Title:     "用户运营"},
		"activity_op": {
			ChannelId: "102",
			Title:     "活动运营"},
		"new_media": {
			ChannelId: "103",
			Title:     "新媒体"},
		"data_op": {
			ChannelId: "104",
			Title:     "数据运营"},
		"video_live": {
			ChannelId: "116",
			Title:     "视频直播"},
		"e_fast_selling": {
			ChannelId: "117",
			Title:     "电商快销"},
		"ASO": {
			ChannelId: "105",
			Title:     "ASO"},
		"SEM": {
			ChannelId: "106",
			Title:     "SEM"},
		"info_stream": {
			ChannelId: "107",
			Title:     "信息流"},
		"marking_promotion": {
			ChannelId: "108",
			Title:     "营销推广"},
		"brand_strategy": {
			ChannelId: "114",
			Title:     "品牌策略"},
		"ad": {
			ChannelId: "115",
			Title:     "创意广告"},
		"create_activity": {
			ChannelId: "121",
			Title:     "创作活动"},
		"career_growth": {
			ChannelId: "110",
			Title:     "职场成长"},
		"product_design": {
			ChannelId: "118",
			Title:     "产品设计"},
		"eff_tool": {
			ChannelId: "119",
			Title:     "效率工具"},
		"management": {
			ChannelId: "120",
			Title:     "经营管理"},
	}

	return Links
}
