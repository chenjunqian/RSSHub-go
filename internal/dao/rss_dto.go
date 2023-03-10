package dao

type RSSFeed struct {
	Title       string
	Link        string
	RSSLink     string
	Description string
	Author      string
	Created     string
	ImageUrl    string
	Items       []RSSItem
	Tag         []string
}

type RSSItem struct {
	Title       string
	Link        string
	Description string
	Content     string
	Author      string
	Created     string
	Thumbnail   string
}

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

type RssFeedChannel struct {
	Id          string
	Title       string
	ChannelDesc string
	ImageUrl    string
	Link        string
	RssLink     string
	Items       []RssFeedItem
	Count       int
}
