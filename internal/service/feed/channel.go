package feed

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/model"
	"rsshub/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)


func GetAllChannelInfoList(ctx context.Context) (feedChannelList []dao.RSSFeed) {
  
  feedChannelList = make([]dao.RSSFeed, 0)
  var feedChannelModeList []model.RssFeedChannel = make([]model.RssFeedChannel, 0)
	if err := service.GetDatabase().Table("rss_feed_channel rfc").
		Select("rfc.*").
    Order("rfc.title asc").
		Find(&feedChannelModeList).Error; err != nil {
		g.Log().Line().Error(ctx, err)
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
