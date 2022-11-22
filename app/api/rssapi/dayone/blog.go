package dayone

import (
	"context"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetMostRead(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	if value, err := component.GetRedis().Do(ctx, "GET", "DAY_ONE_BLOG"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}
	apiUrl := "https://dayoneapp.com/blog/"
	rssData := dao.RSSFeed{
		Title:    "Day One Blog",
		Link:     apiUrl,
		Tag:      []string{"其他"},
		ImageUrl: "https://i0.wp.com/dayoneapp.com/wp-content/uploads/2021/11/favicon-32x32-1.png?fit=32%2C32&ssl=1",
	}
	if resp := component.GetContent(ctx, apiUrl); resp != "" {
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
			rssItem.Content = getFullContent(ctx, rssItem.Link)
			rssItems = append(rssItems, rssItem)
		}

		rssData.Items = rssItems
	}
	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	component.GetRedis().Do(ctx, "SET", "DAY_ONE_BLOG", rssStr)
	component.GetRedis().Do(ctx, "EXPIRE", "DAY_ONE_BLOG", 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}

func getFullContent(ctx context.Context, url string) (content string) {
	if resp := component.GetContent(ctx, url); resp != "" {
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
