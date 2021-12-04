package feed

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/ghash"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gorilla/feeds"
	"github.com/olivere/elastic/v7"
	"rsshub/app/component"
	"rsshub/app/model"
	"strconv"
)

func AddFeedChannelAndItem(feed *feeds.Feed, tagList []string, rsshubLink string) error {

	feedID := strconv.FormatUint(ghash.RSHash64([]byte(feed.Link.Href+feed.Title)), 32)
	feedChannelModel := model.RssFeedChannel{
		Id:          feedID,
		Title:       feed.Title,
		ChannelDesc: feed.Description,
		ImageUrl:    feed.Image.Url,
		Link:        feed.Link.Href,
		RsshubLink:  rsshubLink,
	}

	feedItemModeList := make([]model.RssFeedItem, 0)
	for _, item := range feed.Items {
		feedItem := model.RssFeedItem{
			ChannelId:   feedID,
			Title:       item.Title,
			ChannelDesc: item.Description,
			Link:        item.Link.Href,
			Date:        gtime.New(item.Created.String()),
			Author:      item.Author.Name,
			InputDate:   gtime.Now(),
			Thumbnail:   item.Enclosure.Url,
		}
		uniString := feedItem.Link + feedItem.Title
		feedItemID := strconv.FormatUint(ghash.RSHash64([]byte(uniString)), 32)
		feedItem.Id = feedItemID
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

		_, _ = tx.Save("rss_feed_channel", feedChannelModel)
		_, _ = tx.BatchInsertIgnore("rss_feed_tag", tagModeList)
		_, err = tx.BatchInsertIgnore("rss_feed_item", feedItemModeList)

		return err
	})
	if err != nil {
		g.Log().Error("insert rss feed data failed : ", err)
	}

	bulkRequest := component.GetESClient().Bulk()
	for _, feedItem := range feedItemModeList {
		esFeedItem := model.RssFeedItemESData{
			Id:              feedItem.Id,
			ChannelId:       feedItem.ChannelId,
			Title:           feedItem.Title,
			ChannelDesc:     feedItem.ChannelDesc,
			Thumbnail:       feedItem.Thumbnail,
			Link:            feedItem.Link,
			Date:            feedItem.Date,
			Author:          feedItem.Author,
			InputDate:       feedItem.InputDate,
			ChannelImageUrl: feedChannelModel.ImageUrl,
			ChannelTitle:    feedChannelModel.Title,
			ChannelLink:     feedChannelModel.Link,
		}
		indexReq := elastic.NewBulkIndexRequest().Index("rss_item").Id(feedItem.Id).Doc(esFeedItem)
		bulkRequest.Add(indexReq)
	}
	resp, err := bulkRequest.Do(component.GetESContext())
	if err != nil || resp.Errors {
		respStr := gjson.New(resp)
		g.Log().Errorf("bulk index request failed\nError message : %s \nResponse : %s", err, respStr)
	}

	return err
}
