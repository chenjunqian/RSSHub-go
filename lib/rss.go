package lib

import (
	"time"

	"github.com/gogf/gf/util/gconv"
	"github.com/gorilla/feeds"
)

func GenerateRSS(data map[string]interface{}) string {
	feed := &feeds.Feed{
		Title:       gconv.String(data["title"]),
		Link:        &feeds.Link{Href: gconv.String(data["link"])},
		Description: gconv.String(data["description"]),
		Author:      &feeds.Author{Name: gconv.String(data["author"])},
		Created:     gconv.Time(data["created"]),
	}

	feed.Items = make([]*feeds.Item, 0)
	itemList := data["items"].([]map[string]string)

	for _, item := range itemList {
		createdTime, _ := time.Parse(time.RFC3339, item["created"])
		feedItem := feeds.Item{
			Title:       item["title"],
			Link:        &feeds.Link{Href: item["link"]},
			Description: item["description"],
			Author:      &feeds.Author{Name: item["author"]},
			Created:     createdTime,
		}
		feed.Items = append(feed.Items, &feedItem)
	}

	if result, err := feed.ToRss(); err == nil {
		return result
	} else {
		return ""
	}
}
