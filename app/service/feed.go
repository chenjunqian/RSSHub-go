package service

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/ghash"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gorilla/feeds"
	"rsshub/app/model"
	"strconv"
)

func AddFeedChannelAndItem(feed *feeds.Feed, tagList []string) error {

	feedID := strconv.FormatUint(ghash.RSHash64([]byte(feed.Link.Href+feed.Title)), 32)
	feedChannelModel := model.RssFeedChannel{
		Id:          feedID,
		Title:       feed.Title,
		ChannelDesc: feed.Description,
		ImageUrl:    feed.Image.Url,
		Link:        feed.Link.Href,
	}

	feedItemModeList := make([]model.RssFeedItem, 0)
	for _, item := range feed.Items {
		itemID := strconv.FormatUint(ghash.RSHash64([]byte(item.Link.Href+item.Title)), 32)
		feedItem := model.RssFeedItem{
			Id:          itemID,
			ChannelId:   feedID,
			Title:       item.Title,
			ChannelDesc: item.Description,
			Link:        item.Link.Href,
			Date:        gtime.New(item.Created.String()),
			Author:      item.Author.Name,
			InputDate:   gtime.Now(),
		}
		feedItemModeList = append(feedItemModeList, feedItem)
	}

	tagModeList := make([]model.RssFeedTag, 0)
	for _, tagStr := range tagList {
		if tagStr == "" {
			continue
		}
		tagModel := model.RssFeedTag{
			Name:      tagStr,
			ChannelId: feedID,
			Title:     feed.Title,
			Date:      gtime.Now(),
		}

		tagModeList = append(tagModeList, tagModel)
	}

	err := g.DB().Transaction(func(tx *gdb.TX) error {
		var err error

		_, err = tx.InsertIgnore("rss_feed_channel", feedChannelModel)
		if err != nil {
			return err
		}
		_, err = tx.BatchInsertIgnore("rss_feed_item", feedItemModeList)
		if err != nil {
			return err
		}
		_, err = tx.BatchInsertIgnore("rss_feed_tag", tagModeList)
		if err != nil {
			return err
		}

		return err
	})

	return err
}
