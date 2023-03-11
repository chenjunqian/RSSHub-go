package feed

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/model"
	"rsshub/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetFeedItemByItemId(ctx context.Context, itemId string) (item dao.RssFeedItem) {
	var (
		err           error
		feedItemModel model.RssFeedItem
	)
	if err := service.GetDatabase().Table("rss_feed_item rfi").
		Joins("left join rss_feed_channel rfc on rfi.channel_id=rfc.id").
		Select("rfi.*, rfc.rss_link as rssLink, rfc.title as channelTitle, rfc.image_url as channelImageUrl").
		Where("rfi.id", itemId).
		Find(&feedItemModel).Error; err != nil {
		g.Log().Error(ctx, err)
	}

	err = gconv.Struct(feedItemModel, &item)
	if err != nil {
		g.Log().Error(ctx, err)
	}

	if feedItemModel.Date == nil {
		item.Date = feedItemModel.InputDate.Format("Y-m-d")
	} else {
		item.Date = feedItemModel.Date.Format("Y-m-d")
	}

	if item.Thumbnail == "" {
		item.HasThumbnail = false
	} else {
		item.HasThumbnail = true
	}

	return
}

func GetLatestFeedItem(ctx context.Context, start, size int) (rssFeedItemDtoList []dao.RssFeedItem) {
	var (
		itemList []model.RssFeedItem
	)
	if err := service.GetDatabase().Table("rss_feed_item rfi").
		Select(model.RFIWithoutContentFieldSql + ", rfc.rss_link as rssLink, rfc.title as channelTitle, rfc.image_url as channelImageUrl").
		Joins("inner join rss_feed_channel rfc on rfi.channel_id=rfc.id").
		Group("rfc.id").
		Order("rfi.input_date desc").
		Limit(size).
		Offset(start).
		Find(&itemList).Error; err != nil {
		g.Log().Error(ctx, err)
	}

	rssFeedItemDtoList = make([]dao.RssFeedItem, 0)

	for i := 0; i < len(itemList); i++ {
		var (
			rssFeedItemDto dao.RssFeedItem
			item           model.RssFeedItem
		)
		item = itemList[i]
		rssFeedItemDto = dao.RssFeedItem{
			Id:              item.Id,
			ChannelId:       item.ChannelId,
			Title:           item.Title,
			Description:     item.Description,
			Content:         item.Content,
			Link:            item.Link,
			RssLink:         item.RssLink,
			Author:          item.Author,
			Thumbnail:       item.Thumbnail,
			ChannelImageUrl: item.ChannelImageUrl,
			ChannelTitle:    item.ChannelTitle,
			Count:           1000,
		}
		if item.Date == nil {
			rssFeedItemDto.Date = item.InputDate.Format("Y-m-d")
		} else {
			rssFeedItemDto.Date = item.Date.Format("Y-m-d")
		}

		if item.Content == "" {
			rssFeedItemDto.Content = item.Description
		}

		if rssFeedItemDto.Thumbnail == "" {
			rssFeedItemDto.HasThumbnail = false
		} else {
			rssFeedItemDto.HasThumbnail = true
		}

		rssFeedItemDtoList = append(rssFeedItemDtoList, rssFeedItemDto)
	}

	return
}

func SearchFeedItem(ctx context.Context, keyword string, start, size int) (items []dao.RssFeedItem) {
	var (
		queryString       string
		feedItemModelList []model.RssFeedItem
		count             int64
	)

	if size == 0 {
		size = 10
	}

	if start <= 0 {
		start = 0
	} else {
		start = (start - 1) * size
	}

	queryString = "(SELECT id FROM rss_feed_item fts WHERE MATCH (title, content, author,description) AGAINST ('" + keyword + "'))"
	if err := service.GetDatabase().Table("rss_feed_item rfi").
		Select(model.RFIWithoutContentFieldSql + ", rfc.rss_link as rssLink, rfc.title as channelTitle, rfc.image_url as channelImageUrl").
		Joins("inner join rss_feed_channel rfc on rfi.channel_id=rfc.id and rfi.id in " + queryString).
		Order("rfi.input_date desc").
		Limit(size).
		Offset(start).
		Find(&feedItemModelList).Error; err != nil {
		g.Log().Error(ctx, err)
	}

	// Get all count
	var queryCountString = "MATCH (title, content, author,description) AGAINST ('" + keyword + "')"
	if result := service.GetDatabase().Table("rss_feed_item rfi").
		Where(queryCountString).
		Count(&count); result.Error != nil {
		g.Log().Error(ctx, result.Error)
	}

	for _, itemModel := range feedItemModelList {
		var (
			feedItemDao dao.RssFeedItem
		)

		gconv.Struct(itemModel, &feedItemDao)
		feedItemDao.Count = int(count)
		if feedItemDao.Thumbnail != "" {
			feedItemDao.HasThumbnail = true
		}
		if itemModel.Date == nil {
			feedItemDao.Date = itemModel.InputDate.Format("Y-m-d")
		} else {
			feedItemDao.Date = itemModel.Date.Format("Y-m-d")
		}

		items = append(items, feedItemDao)
	}

	return
}
