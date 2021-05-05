package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"rsshub/app/model/biz"
)

func GetFeedItemByTag(start, size int, channelId string) (itemList []biz.RssFeedItemData) {

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
