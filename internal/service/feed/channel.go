package feed

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/model"
	"rsshub/internal/model/dto"
	"rsshub/internal/service"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mmcdole/gofeed"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllChannelInfoList(ctx context.Context) (feedChannelList []dao.RSSFeed) {

	feedChannelList = make([]dao.RSSFeed, 0)
	var feedChannelModeList []model.RssFeedChannel = make([]model.RssFeedChannel, 0)
	if service.GetDatabase() == nil {
		return
	}
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

func GetAllDefinedRouters(ctx context.Context) (routerDataList []model.RouterInfoData) {
	var (
		routerArray []ghttp.RouterItem
	)
	routerArray = g.Server().GetRoutes()
	routerDataList = make([]model.RouterInfoData, 0)
	if len(routerArray) > 0 {
		for _, router := range routerArray {
			if strings.Contains(router.Route, ":") || strings.Contains(router.Route, "*") || strings.HasPrefix(router.Route, "/api") {
				continue
			}
			routerInfoData := model.RouterInfoData{
				Route: router.Route,
				Port:  router.Address,
			}
			routerDataList = append(routerDataList, routerInfoData)
		}
	}
	return
}

func AddFeedChannelAndItem(ctx context.Context, feed *gofeed.Feed, rsshubLink string) error {
	var (
		err              error
		feedChannelModeL model.RssFeedChannel
		feedItemModeList []model.RssFeedItem
	)
	feedChannelModeL, feedItemModeList = assembleFeedChannlAndItem(ctx, feed, rsshubLink)
	err = storeFeedChannelAndItems(ctx, feedChannelModeL, feedItemModeList)
	return err
}

func assembleFeedChannlAndItem(ctx context.Context, feed *gofeed.Feed, rsshubLink string) (feedChannelModel model.RssFeedChannel, feedItemModeList []model.RssFeedItem) {

	feedID := strconv.FormatUint(ghash.RS64([]byte(feed.Link+feed.Title)), 32)
	feedChannelModel = model.RssFeedChannel{
		Id:          feedID,
		Title:       feed.Title,
		ChannelDesc: feed.Description,
		Link:        feed.Link,
		RssLink:     rsshubLink,
	}
	if feed.Image != nil {
		feedChannelModel.ImageUrl = feed.Image.URL
	}

	feedItemModeList = make([]model.RssFeedItem, 0)
	for _, item := range feed.Items {
		var (
			thumbnail string
			author    string
		)
		if len(item.Enclosures) > 0 {
			thumbnail = item.Enclosures[0].URL
		}

		if thumbnail == "" {
			thumbnail = getDescriptionThumbnail(item.Description)
		}

		if thumbnail == "" {
			thumbnail = getDescriptionThumbnail(item.Content)
		}

		if len(item.Authors) > 0 {
			author = item.Authors[0].Name
		}

		feedItem := model.RssFeedItem{
			ChannelId:   feedID,
			Title:       item.Title,
			Description: item.Description,
			Content:     item.Content,
			Link:        item.Link,
			Date:        gtime.New(item.Published),
			Author:      author,
			InputDate:   gtime.Now(),
			Thumbnail:   thumbnail,
		}
		uniString := feedItem.Link + feedItem.Title
		feedItemID := strconv.FormatUint(ghash.RS64([]byte(uniString)), 32)
		feedItem.Id = feedItemID
		feedItemModeList = append(feedItemModeList, feedItem)
	}

	return
}

func storeFeedChannelAndItems(ctx context.Context, feedChannelModel model.RssFeedChannel, feedItemModeList []model.RssFeedItem) error {
	if service.GetDatabase() == nil {
		return nil
	}
	err := service.GetDatabase().Transaction(func(tx *gorm.DB) error {
		var err error

		err = tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&feedChannelModel).Error

		if err != nil {
			g.Log().Error(ctx, "inser feedChannelModel failed : ", err, " ,feedChannelModel is ", gjson.MustEncode(feedChannelModel))
			return err
		}

		err = tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&feedItemModeList).Error

		if err != nil {
			g.Log().Error(ctx, "inser feedItemModeList failed : ", err, " ,feedItemModeList is ", gjson.MustEncode(feedItemModeList))
			return err
		}

		return err
	})
	if err != nil {
		g.Log().Error(ctx, "insert rss feed data failed : ", gjson.MustEncode(err))
	}

	return err
}

func getDescriptionThumbnail(htmlStr string) (thumbnail string) {

	docs := soup.HTMLParse(htmlStr)
	firstImgDoc := docs.Find("img")
	if firstImgDoc.Pointer != nil {
		thumbnail = firstImgDoc.Attrs()["src"]
	}

	return
}

func GetLatestFeedItem(ctx context.Context, start, size int) (rssFeedItemDtoList []dto.RssFeedItem) {
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

	rssFeedItemDtoList = make([]dto.RssFeedItem, 0)

	for i := 0; i < len(itemList); i++ {
		var (
			rssFeedItemDto dto.RssFeedItem
			item           model.RssFeedItem
		)
		item = itemList[i]
		rssFeedItemDto = dto.RssFeedItem{
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
