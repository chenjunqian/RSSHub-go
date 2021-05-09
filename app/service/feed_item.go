package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"rsshub/app/model/biz"
)

func GetFeedItemByChannelId(start, size int, channelId string) (itemList []biz.RssFeedItemData) {

	if err := g.DB().Table("rss_feed_item rfi").LeftJoin("rss_feed_channel rfc", "rfi.channel_id=rfc.id").
		Fields("rfi.*, rfc.title as channelTitle, rfc.image_url as channelImageUrl").
		Where("rfi.channel_id", channelId).
		Order("rfi.input_date desc").
		Limit(start, size).
		Structs(&itemList); err != nil {
		glog.Line().Error(err)
	}

	return
}

func GetFeedItemListByUserId(userId string, start, size int) (itemList []biz.RssFeedItemData) {
	if err := g.DB().Table("rss_feed_item rfi").
		InnerJoin("user_feed_channel_sub ufcs", "ufcs.channel_id=rfc.channel_id").
		InnerJoin("rss_feed_channel rfc", "ufcs.channel_id=rfc.id").
		Fields("rfi.*, rfc.title as channelTitle, rfc.image_url as channelImageUrl").
		Where("ufcs.user_id", userId).
		Order("rfi.input_date desc").
		Limit(start, size).
		Structs(&itemList); err != nil {
		glog.Line().Error(err)
	}

	return
}
