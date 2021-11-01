package model

import "github.com/gogf/gf/os/gtime"

type RssFeedItemESData struct {
	Id              string      `json:"id"`
	ChannelId       string      `json:"channel_id"`
	Title           string      `json:"title"`
	ChannelDesc     string      `json:"channel_desc"`
	Link            string      `json:"link"`
	Date            *gtime.Time `json:"date"`
	Author          string      `json:"author"`
	InputDate       *gtime.Time `json:"input_date"`
	ChannelImageUrl string      `json:"channel_image_link"`
	ChannelTitle    string      `json:"channel_title"`
	ChannelLink     string      `json:"channel_link"`
}
