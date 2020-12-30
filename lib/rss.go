package lib

import (
	"rsshub/app/dao"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gorilla/feeds"
)

func GenerateRSS(data dao.RSSFeed) string {
	feed := &feeds.Feed{
		Title:       data.Title,
		Link:        &feeds.Link{Href: data.Link},
		Description: data.Description,
		Author:      &feeds.Author{Name: data.Author},
		Created:     gconv.Time(data.Created),
	}

	feed.Items = make([]*feeds.Item, 0)
	itemList := data.Items

	for _, item := range itemList {
		feedItem := feeds.Item{
			Title:       item.Title,
			Link:        &feeds.Link{Href: item.Link},
			Description: item.Description,
			Author:      &feeds.Author{Name: item.Author},
		}
		if item.Created != "" {
			feedItem.Created = gtime.NewFromStr(item.Created).Time
		}
		feed.Items = append(feed.Items, &feedItem)
	}

	if result, err := feed.ToRss(); err == nil {
		return result
	} else {
		return ""
	}
}
