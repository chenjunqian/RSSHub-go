package dx2025

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
)

type Controller struct {
}

type IndustryInfoRouteConfig struct {
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
	respDocs := soup.HTMLParse(htmlStr)
	dataDocsList := respDocs.FindAll("div", "class", "post-wrapper-hentry")
	if len(dataDocsList) > 15 {
		dataDocsList = dataDocsList[:15]
	}
	for _, dataDocs := range dataDocsList {
		var imageLink string
		var title string
		var link string
		var content string
		var time string
		imageWrap := dataDocs.Find("figure", "class", "post-thumbnail")
		if imageWrap.Error == nil {
			imageLink = imageWrap.Find("img").Attrs()["src"]
		}
		timeWrap := dataDocs.Find("time", "class", "updated")
		if timeWrap.Error == nil {
			time = timeWrap.Attrs()["datetime"]
		}
		titleWrap := dataDocs.Find("h1", "class", "entry-title")
		if titleWrap.Error == nil {
			title = titleWrap.Find("a").Text()
			link = titleWrap.Find("a").Attrs()["href"]
		}

		content = parseCommonDetail(link)

		rssItem := dao.RSSItem{
			Title:       title,
			Link:        link,
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
		respString = resp.ReadAllString()
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("div", "class", "post-wrapper-hentry")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request dx2025 article detail failed, link  %s \nerror : %s", detailLink, err)
	}

	return
}

func getIndustryInfoLinks() map[string]IndustryInfoRouteConfig {
	Links := map[string]IndustryInfoRouteConfig{
		"new-tch-info": {
			ChannelId: "next-generation-information-technology",
			Title:     "新一代信息技术"},
		"robot-idt-info": {
			ChannelId: "high-grade-cnc-machine-tools-and-robots",
			Title:     "高档数控机床和机器人"},
		"aerospace-info": {
			ChannelId: "aerospace-equipment",
			Title:     "航空航天装备"},
		"marine-info": {
			ChannelId: "marine-engineering-equipment-and-high-tech-ships",
			Title:     "海工装备及高技术船舶"},
		"transportation-info": {
			ChannelId: "advanced-rail-transportation-equipment",
			Title:     "先进轨道交通装备"},
		"energy-vehicles-info": {
			ChannelId: "energy-saving-and-new-energy-vehicles",
			Title:     "节能与新能源汽车"},
		"electric-equipment-info": {
			ChannelId: "electric-equipment",
			Title:     "电力装备"},
		"agricultural-equipment-info": {
			ChannelId: "agricultural-machinery-equipment",
			Title:     "农机装备"},
		"material-info": {
			ChannelId: "new-material",
			Title:     "新材料"},
		"biomedicine-medical-info": {
			ChannelId: "biomedicine-and-medical-devices",
			Title:     "生物医药及医疗器械"},
		"modern-service-info": {
			ChannelId: "modern-service-industry",
			Title:     "现代服务业"},
		"manufacturing-info": {
			ChannelId: "manufacturing-talent",
			Title:     "制造业人才"},
	}

	return Links
}
