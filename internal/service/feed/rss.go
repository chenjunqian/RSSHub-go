package feed

import (
	"rsshub/internal/dao"
	"strings"

	"github.com/GuoShaoOrg/feeds"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func GenerateRSS(data dao.RSSFeed, rsshubLink string) string {

	feed := &feeds.Feed{
		Title:       data.Title,
		Link:        &feeds.Link{Href: data.Link},
		Description: data.Description,
		Author:      &feeds.Author{Name: data.Author},
		Image:       &feeds.Image{Link: data.ImageUrl, Url: data.ImageUrl},
		Created:     gconv.Time(data.Created),
	}

	feed.Items = make([]*feeds.Item, 0)

	for _, item := range data.Items {
		feedItem := feeds.Item{
			Title:       item.Title,
			Link:        &feeds.Link{Href: item.Link},
			Description: item.Description,
			Content:     item.Content,
			Author:      &feeds.Author{Name: item.Author},
			Enclosure:   &feeds.Enclosure{Url: item.Thumbnail},
		}
		if item.Created != "" {
			if gtime.NewFromStr(item.Created) != nil {
				feedItem.Created = gtime.NewFromStr(item.Created).Time
			}
		} else {
			nowTime := gtime.Now().Format("Y-m-d H:i:s.u")
			if gtime.NewFromStr(nowTime) != nil {
				feedItem.Created = gtime.NewFromStr(nowTime).Time
			}
		}
		feed.Items = append(feed.Items, &feedItem)
	}

	if result, err := feed.ToAtom(); err == nil {
		return result
	} else {
		return ""
	}
}

func GenerateContent(content string) (description string) {
	var imageHtml string
	var contentHtml string
	var htmlString string
	var imageLink string
	if strings.HasPrefix(content, "http://") || strings.HasPrefix(content, "https://") {
		imageLink = content
	}
	if imageLink != "" {
		imageHtml = "<img src=" + imageLink + " style='width:100%' >"
		htmlString = "<meta name='referrer' content='no-referrer' /><div style='position: relative;text-align: left;'>" + imageHtml + "</div>"
	} else {
		contentHtml = "<div >" + content + "</div>"
		htmlString = "<div style='position: relative;text-align: left;'>" + contentHtml + "</div>"
	}
	description = htmlString
	return
}
