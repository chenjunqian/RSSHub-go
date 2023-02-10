package model

type RssFeedChannel struct {
	Id          string `gorm:"column:id;primaryKey"   json:"id"`
	Title       string `gorm:"column:title"        json:"title"`
	ChannelDesc string `gorm:"column:channel_desc" json:"channelDesc"`
	ImageUrl    string `gorm:"column:image_url"    json:"imageUrl"`
	Link        string `gorm:"column:link"         json:"link"`
	RssLink     string `gorm:"column:rss_link"  json:"rssLink"`
}

func (RssFeedChannel) TableName() string {
	return "rss_feed_channel"
}
