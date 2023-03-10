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
