package feed

import (
	"context"
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/model"
)


func GetAllChannelInfoList(ctx context.Context) (feedChannelList []dao.RSSFeed) {
  
  feedChannelList = make([]dao.RSSFeed, 0)
  var feedChannelModeList []model.RssFeedChannel = make([]model.RssFeedChannel, 0)
	if err := component.GetDatabase().Table("rss_feed_channel rfc").
		Select("rfc.*").
    Order("rfc.title asc").
		Find(&feedChannelModeList).Error; err != nil {
		component.GetLogger().Error(ctx, err)
	} else {
    for _, item := range feedChannelModeList {
      var feedChannelInfo = dao.RSSFeed{}
      feedChannelInfo.Title = item.Title
      feedChannelInfo.Description = item.ChannelDesc
      feedChannelInfo.ImageUrl = item.ImageUrl
      feedChannelInfo.Link = item.ImageUrl
      feedChannelInfo.RSSLink = item.RssLink
      feedChannelList = append(feedChannelList, feedChannelInfo)
    }
  }

	return
}
