package webapi

import (
	"rsshub/internal/dao"
	feedService "rsshub/internal/service/feed"
	response "rsshub/internal/service/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

func (ctl *Controller) IndexTpl(req *ghttp.Request) {
	var (
		latestFeedItemList []dao.RssFeedItem
	)

	latestFeedItemList = feedService.GetLatestFeedItem(req.Context(), 0, 10)

	req.Response.WriteTpl("home.html", g.Map{
		"name":               "RSS Go",
		"latestFeedItemList": latestFeedItemList,
	})
}

func (ctl *Controller) FeedChannelDetail(req *ghttp.Request) {
	var (
		channelInfo dao.RssFeedChannel
		channelId   string
	)

	channelId = req.Get("id").String()
	channelInfo = feedService.GetChannelInfoByChannelId(req.Context(), channelId)
	req.Response.WriteTpl("channel.html", g.Map{
		"name":        "RSS Go",
		"channelInfo": channelInfo,
	})
}

func (ctl *Controller) GetAllRssResource(req *ghttp.Request) {

	routerDataList := feedService.GetAllDefinedRouters(gctx.New())
	response.JsonExit(req, 0, "success", routerDataList)
}

func (ctl *Controller) GetAllFeedChannelInfoList(req *ghttp.Request) {

	var feedChannelInfoList []dao.RSSFeed

	feedChannelInfoList = feedService.GetAllChannelInfoList(req.Context())

	response.JsonExit(req, 0, "success", feedChannelInfoList)
}
