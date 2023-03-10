package model

import "github.com/gogf/gf/v2/os/gtime"

type RssFeedItem struct {
	Id              string      `gorm:"column:id"`
	ChannelId       string      `gorm:"column:channel_id"`
	Title           string      `gorm:"column:title"`
	Description     string      `gorm:"column:description"`
	Link            string      `gorm:"column:link"`
	RssLink         string      `gorm:"column:rssLink"`
	Date            *gtime.Time `gorm:"column:date"`
	Author          string      `gorm:"column:author"`
	InputDate       *gtime.Time `gorm:"column:input_date"`
	Thumbnail       string      `gorm:"column:thumbnail"`
	Content         string      `gorm:"column:content"`
	ChannelImageUrl string      `gorm:"column:channelImageUrl"`
	ChannelTitle    string      `gorm:"column:channelTitle"`
}

func (RssFeedItem) TableName() string {
	return "rss_feed_item"
}

var (
	RFIWithoutContentFieldSql = "rfi.id, rfi.channel_id, rfi.title, rfi.description, rfi.content, rfi.link, rfi.date, rfi.author, rfi.input_date, rfi.thumbnail"
)
