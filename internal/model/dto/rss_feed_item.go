package dto

type RssFeedItem struct {
	Id              string
	ChannelId       string
	Title           string
	Description     string
	Link            string
	RssLink         string
	Date            string
	Author          string
	Thumbnail       string
	Content         string
	ChannelImageUrl string
	ChannelTitle    string
	Count           int
	HasThumbnail    bool
}
