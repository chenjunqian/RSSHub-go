package ifeng

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"regexp"
	"rsshub/app/dao"
)

type Controller struct {
}

type NewsRouteConfig struct {
	ChannelId string
	Title     string
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func commonParser(htmlStr string) (items []dao.RSSItem) {
	reg := regexp.MustCompile(`var allData = (.*?);\n`)
	jsonStr := reg.FindStringSubmatch(htmlStr)
	if len(jsonStr) <= 1 {
		return items
	}
	respJson := gjson.New(jsonStr[1])
	dataJsonList := respJson.GetJsons("newsstream")
	if len(dataJsonList) > 20 {
		dataJsonList = dataJsonList[:20]
	}
	for _, dataJson := range dataJsonList {
		var imageLink string
		var title string
		var link string
		var content string
		var time string
		imageJson := dataJson.GetJson("thumbnails.image")
		if imageJson.IsNil() {
			imageLink = dataJson.GetString("thumbnail")
		} else {
			imageLink = dataJson.GetString("thumbnails.image.0.url")
		}
		title = dataJson.GetString("title")
		content = dataJson.GetString("summary")
		time = dataJson.GetString("newsTime")
		link = dataJson.GetString("url")
		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Description: fmt.Sprintf("<img src='%s'><br>%s", imageLink, content),
			Created:     time,
		}
		items = append(items, rssItem)
	}
	return
}

func moneyCommonParser(htmlStr, typeStr string) (items []dao.RSSItem) {
	reg := regexp.MustCompile(`var allData = (.*?);\n`)
	jsonStr := reg.FindStringSubmatch(htmlStr)
	if len(jsonStr) <= 1 {
		return items
	}
	respJson := gjson.New(jsonStr[1])
	dataJsonList := respJson.GetJsons(typeStr)
	if len(dataJsonList) > 20 {
		dataJsonList = dataJsonList[:20]
	}
	for _, dataJson := range dataJsonList {
		var imageLink string
		var title string
		var author string
		var link string
		var content string
		var time string
		imageJson := dataJson.GetJson("thumbnails.image")
		if imageJson.IsNil() {
			imageLink = dataJson.GetString("thumbnail")
		} else {
			imageLink = dataJson.GetString("thumbnails.image.0.url")
		}
		title = dataJson.GetString("title")
		content = dataJson.GetString("summary")
		time = dataJson.GetString("newsTime")
		link = dataJson.GetString("url")
		author = dataJson.GetString("editorName")
		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
			Author:      author,
			Description: fmt.Sprintf("<img src='%s'><br>%s", imageLink, content),
			Created:     time,
		}
		items = append(items, rssItem)
	}
	return
}

func getNewsInfoLinks() map[string]NewsRouteConfig {
	Links := map[string]NewsRouteConfig{
		"xijinping": {
			ChannelId: "xijinping",
			Title:     "新时代"},
		"living": {
			ChannelId: "living",
			Title:     "在人间"},
		"bigfish": {
			ChannelId: "bigfish",
			Title:     "大鱼漫画"},
		"globalyouth": {
			ChannelId: "globalyouth",
			Title:     "地球青年"},
		"warmstory": {
			ChannelId: "warmstory",
			Title:     "暖新闻"},
		"fhphotos": {
			ChannelId: "fhphotos",
			Title:     "图片特刊"},
	}

	return Links
}

func getFinanceInfoLinks() map[string]NewsRouteConfig {
	Links := map[string]NewsRouteConfig{
		"hk_hotpoint": {
			ChannelId: "shanklist/1-69-35252-",
			Title:     "港股 - 热点"},
		"hk_baishitong": {
			ChannelId: "shanklist/1-69-35256-",
			Title:     "港股 - 百事通"},
		"hk_xiaozhishi": {
			ChannelId: "shanklist/1-69-35255-",
			Title:     "港股 - 小知识"},
		"hk_zhuanlan": {
			ChannelId: "shanklist/1-69-35258-",
			Title:     "港股 - 深度专栏"},
		"hk_jigou_dongtai": {
			ChannelId: "shanklist/1-69-35253-",
			Title:     "港股 - 机构动态"},
		"hk_tianxia": {
			ChannelId: "shanklist/1-69-35260-",
			Title:     "港股 - 天下财经智库"},
		"ipo_kechuang": {
			ChannelId: "shanklist/1-95385-",
			Title:     "新股 - 科创板"},
	}

	return Links
}

func getMoneyInfoLinks() map[string]NewsRouteConfig {
	Links := map[string]NewsRouteConfig{
		"hot": {
			ChannelId: "hotStream",
			Title:     "理财 - 热点"},
		"bank": {
			ChannelId: "bankStream",
			Title:     "理财 - 银行"},
		"insure": {
			ChannelId: "insureStream",
			Title:     "理财 - 保险"},
		"fund": {
			ChannelId: "fundStream",
			Title:     "理财 - 基金"},
	}

	return Links
}

func getEntertainmentInfoLinks() map[string]NewsRouteConfig {
	Links := map[string]NewsRouteConfig{
		"star": {
			ChannelId: "star",
			Title:     "娱乐 - 明星"},
		"movie": {
			ChannelId: "bankStream",
			Title:     "娱乐 - 电影"},
		"tv": {
			ChannelId: "insureStream",
			Title:     "娱乐 - 电视"},
		"music": {
			ChannelId: "fundStream",
			Title:     "娱乐 - 音乐"},
	}

	return Links
}

func getCultureInfoLinks() map[string]NewsRouteConfig {
	Links := map[string]NewsRouteConfig{
		"read": {
			ChannelId: "shanklist/17-35108-",
			Title:     "文化 - 读书"},
		"artist": {
			ChannelId: "shanklist/17-35106-",
			Title:     "文化 - 文艺家"},
		"insight": {
			ChannelId: "shanklist/17-35105-",
			Title:     "文化 - 洞见"},
		"news": {
			ChannelId: "shanklist/17-35104-",
			Title:     "文化 - 资讯"},
	}

	return Links
}

func getFashionInfoLinks() map[string]NewsRouteConfig {
	Links := map[string]NewsRouteConfig{
		"trends": {
			ChannelId: "trends",
			Title:     "时尚 - 时装"},
		"beauty": {
			ChannelId: "beauty",
			Title:     "时尚 - 美容美体"},
		"lifestyle": {
			ChannelId: "lifestyle",
			Title:     "时尚 - 生活"},
		"emotion": {
			ChannelId: "emotion",
			Title:     "时尚 - 情感"},
	}

	return Links
}

func getAutoInfoLinks() map[string]NewsRouteConfig {
	Links := map[string]NewsRouteConfig{
		"xinche": {
			ChannelId: "xinche",
			Title:     "汽车 - 新车"},
		"shijia": {
			ChannelId: "shijia",
			Title:     "汽车 - 试驾"},
		"daogou": {
			ChannelId: "daogou",
			Title:     "汽车 - 导购"},
		"hangye": {
			ChannelId: "hangye",
			Title:     "汽车 - 行业"},
	}

	return Links
}

func getTechInfoLinks() map[string]NewsRouteConfig {
	Links := map[string]NewsRouteConfig{
		"index": {
			ChannelId: "",
			Title:     "科技 - 资讯"},
		"digi": {
			ChannelId: "digi",
			Title:     "科技 - 试驾"},
		"mobile": {
			ChannelId: "mobile",
			Title:     "科技 - 手机"},
		"hangye": {
			ChannelId: "hangye",
			Title:     "科技 - 行业"},
	}

	return Links
}
