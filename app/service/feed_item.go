package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"rsshub/app/model/biz"
)

func GetFeedItemByTag(start, size int, channelId string) (itemList []biz.RssFeedItemData) {

	if err := g.DB().Table("rss_feed_item").
		Fields("*").
		Where("channel_id", channelId).
		Limit(start, size).
		Structs(&itemList); err != nil {
		glog.Line().Error(err)
	}

	return
}
