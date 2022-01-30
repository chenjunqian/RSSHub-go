package dayone

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetMostRead(req *ghttp.Request) {

	if value, err := g.Redis().DoVar("GET", "DAY_ONE_BLOG"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://dayoneapp.com/blog/"
	rssData := dao.RSSFeed{
		Title:    "Day One Blog",
		Link:     apiUrl,
		Tag:      []string{"其他"},
		ImageUrl: "https://i0.wp.com/dayoneapp.com/wp-content/uploads/2021/11/favicon-32x32-1.png?fit=32%2C32&ssl=1",
	}
	if resp := component.GetContent(apiUrl); resp != "" {
		docs := soup.HTMLParse(resp)
		blogItemWrapper := docs.Find("div", "class", "container--inner")
		blogItemList := blogItemWrapper.FindAll("div")
		if len(blogItemList) > 15 {
			blogItemList = blogItemList[:15]
		}
		rssItems := make([]dao.RSSItem, 0)
		for _, blogItem := range blogItemList {
			rssItem := dao.RSSItem{}
			title := blogItem.Find("a").Text()
			link := "https://dayoneapp.com" + blogItem.Find("a").Attrs()["href"]
			dateUserStr := blogItem.Find("span").FullText()
			dateUserStrArr := strings.Split(dateUserStr, " by ")
			if len(dateUserStrArr) > 1 {
				rssItem.Created = getParseDate(dateUserStrArr[0])
				rssItem.Author = dateUserStrArr[1]
			} else {
				rssItem.Created = getParseDate(dateUserStrArr[0])
			}

			rssItem.Title = title
			rssItem.Link = link
			rssItem.Description = getFullDescription(rssItem.Link)
			rssItems = append(rssItems, rssItem)
		}

		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "DAY_ONE_BLOG", rssStr)
	g.Redis().DoVar("EXPIRE", "DAY_ONE_BLOG", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}

func getFullDescription(url string) (content string) {
	if resp := component.GetContent(url); resp != "" {
		docs := soup.HTMLParse(resp)
		content = docs.Find("main").HTML()
	}
	return
}

func getParseDate(date string) (formatDate string) {
	dateStr := strings.Split(date, " ")
	monthMap := map[string]string{
		"January":   "01",
		"February":  "02",
		"March":     "03",
		"April":     "04",
		"May":       "05",
		"June":      "06",
		"July":      "07",
		"August":    "08",
		"September": "09",
		"October":   "10",
		"November":  "11",
		"December":  "12",
	}
	if len(dateStr) >= 3 {
		month := monthMap[dateStr[0]]
		day := dateStr[1]
		day = strings.ReplaceAll(day, ",", "")
		if len(day) == 1 {
			day = "0" + day
		}
		year := dateStr[2]
		formatDate = year + "-" + month + "-" + day
		formatDateTime, err := time.Parse("2006-01-02", formatDate)
		if err == nil {
			formatDate = formatDateTime.String()
		}
	}
	return
}
