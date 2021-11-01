package biz

import "github.com/gogf/gf/os/gtime"

type RssFeedItemData struct {
	Id              string
	ChannelId       string
	Title           string
	ChannelDesc     string
	Link            string
	Date            *gtime.Time
	Author          string
	InputDate       *gtime.Time
	RsshubLink      string
	ChannelImageUrl string
	ChannelTitle    string
	Marked          int
	Sub             int
}
