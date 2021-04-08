package webApi

type Controller struct {
}

type RouterInfoData struct {
	Route string
	Port  string
}

type FeedTagReqData struct {
	Start int
	Size  int
}

type FeedChannelReqData struct {
	Start int
	Size  int
	Name  string `v:"required"`
}

type FeedTagByChannelIdReqData struct {
	Start     int
	Size      int
	ChannelId string `p:"channelId" v:"required"`
}
