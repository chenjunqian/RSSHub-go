package duozhi

import (
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"
)

type Controller struct {
}

type IndustryNewsRouteConfig struct {
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
	dataDocsList := respDocs.FindAll("div", "class", "post-item")
	if len(dataDocsList) > 20 {
		dataDocsList = dataDocsList[:20]
	}
	for _, dataDocs := range dataDocsList {
		var (
			imageLink string
			title     string
			link      string
			content   string
			author    string
			time      string
		)

		postImageWrap := dataDocs.Find("a", "class", "post-img")
		if postImageWrap.Error == nil {
			title = postImageWrap.Attrs()["title"]
			link = postImageWrap.Attrs()["href"]
			imageStyleStr := postImageWrap.Attrs()["style"]
			imageStyleStrs := strings.Split(imageStyleStr, "url(")
			if len(imageStyleStrs) >= 2 {
				imageLink = imageStyleStrs[1]
			}
		}

		content = parseCommonDetail(link)

		authorWrap := dataDocs.Find("span", "class", "post-attr")
		if authorWrap.Error == nil {
			author = authorWrap.Text()
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

func parseCommonDetail(detailLink string) (detailData string) {
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
		articleElem = docs.Find("div", "class", "c2")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request duozhi article detail failed, link  %s \nerror : %s", detailLink, err)
	}

	return
}

func getIndustryNewsLinks() map[string]IndustryNewsRouteConfig {
	Links := map[string]IndustryNewsRouteConfig{
		"insight": {
			ChannelId: "insight",
			Tags:      []string{"教育", "其他"},
			Title:     "观察"},
		"preschool": {
			ChannelId: "preschool",
			Tags:      []string{"教育", "早教"},
			Title:     "早教"},
		"K12": {
			ChannelId: "K12",
			Tags:      []string{"教育", "K12"},
			Title:     "K12"},
		"qualityedu": {
			ChannelId: "qualityedu",
			Tags:      []string{"教育"},
			Title:     "素质教育"},
		"adultedu": {
			ChannelId: "adult",
			Tags:      []string{"教育", "职教"},
			Title:     "职教/大学生"},
		"EduInformatization": {
			ChannelId: "EduInformatization",
			Tags:      []string{"教育", "科技"},
			Title:     "信息化教育"},
		"earnings": {
			ChannelId: "earnings",
			Tags:      []string{"教育", "财经"},
			Title:     "财报"},
		"privateschools": {
			ChannelId: "privateschools",
			Tags:      []string{"教育"},
			Title:     "民办教育"},
		"overseas": {
			ChannelId: "overseas",
			Tags:      []string{"教育", "留学"},
			Title:     "留学"},
	}

	return Links
}
