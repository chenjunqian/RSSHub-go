package model

import "github.com/gogf/gf/v2/os/gtime"

type RssFeedItem struct {
	Id              string      `gorm:"column:id,primary"  json:"id"`          //
	ChannelId       string      `gorm:"column:channel_id"  json:"channelId"`   //
	Title           string      `gorm:"column:title"       json:"title"`       //
	Description     string      `gorm:"column:description" json:"description"` //
	Link            string      `gorm:"column:link"        json:"link"`        //
	RssLink         string      `gorm:"column:rssLink"     json:"rssLink"`     //
	Date            *gtime.Time `gorm:"column:date"        json:"date"`        //
	Author          string      `gorm:"column:author"      json:"author"`      //
	InputDate       *gtime.Time `gorm:"column:input_date"  json:"inputDate"`   //
	Thumbnail       string      `gorm:"column:thumbnail"   json:"thumbnail"`   //
	Content         string      `gorm:"column:content"     json:"content"`     //
	ChannelImageUrl string      `gorm:"column:channelImageUrl"`
	ChannelTitle    string      `gorm:"column:channelTitle"`
}

func (RssFeedItem) TableName() string {
	return "rss_feed_item"
}

var (
	RFIWithoutContentFieldSql = "rfi.id, rfi.channel_id, rfi.title, rfi.description, rfi.link, rfi.date, rfi.author, rfi.input_date, rfi.thumbnail"
)
