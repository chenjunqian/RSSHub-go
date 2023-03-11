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

func (ctl *Controller) SearchFeedItems(req *ghttp.Request) {
	var (
		feedItemList  []dao.RssFeedItem
		searchKeyword string
		start         int
		totalPage     int
		count         int
	)

	searchKeyword = req.Get("keyword").String()
	start = req.Get("start").Int()
	feedItemList = feedService.SearchFeedItem(req.Context(), searchKeyword, start, 0)
	if len(feedItemList) > 0 {
		count = feedItemList[0].Count
		totalPage = count / 10
	}
	req.Response.WriteTpl("search.html", g.Map{
		"name":          "RSS Go",
		"searchKeyword": searchKeyword,
		"currentPage":   start,
		"totalPage":     totalPage,
		"feedItems":     feedItemList,
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

func (ctl *Controller) FeedItemDetail(req *ghttp.Request) {
	var (
		itemInfo dao.RssFeedItem
		itemId   string
	)

	itemId = req.Get("id").String()
	itemInfo = feedService.GetFeedItemByItemId(req.Context(), itemId)
	req.Response.WriteTpl("item.html", g.Map{
		"name":     "RSS Go",
		"itemInfo": itemInfo,
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
