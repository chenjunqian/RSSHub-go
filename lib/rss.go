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
		Image:       &feeds.Image{Url: data.ImageUrl},
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
			if gtime.NewFromStr(item.Created) != nil {
				feedItem.Created = gtime.NewFromStr(item.Created).Time
			}
		}
		feed.Items = append(feed.Items, &feedItem)
	}

	if result, err := feed.ToRss(); err == nil {
		return result
	} else {
		return ""
	}
}

func GenerateDescription(imageLink, content string) (description string) {
	var imageHtml string
	if imageLink != "" {
		imageHtml = "<img src=" + imageLink + " style='width:100%' >"
	}
	contentHtml := "<div style='position: absolute;bottom: 8px;left: 8px;' >" + content + "</div>"
	htmlString := "<div style='position: relative;text-align: center;color: white;'>" + imageHtml + contentHtml + "</div>"
	description = htmlString
	return
}
