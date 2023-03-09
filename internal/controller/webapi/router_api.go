package webapi

import (
	"rsshub/internal/dao"
	"rsshub/internal/model/dto"
	feedService "rsshub/internal/service/feed"
	response "rsshub/internal/service/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"golang.org/x/net/context"
)

func (ctl *Controller) IndexTpl(req *ghttp.Request) {
	var (
		latestFeedItemList []dto.RssFeedItem
	)

	latestFeedItemList = feedService.GetLatestFeedItem(gctx.New(), 0, 10)

	req.Response.WriteTpl("home.html", g.Map{
		"name":               "RSS Go",
		"latestFeedItemList": latestFeedItemList,
	})
}

func (ctl *Controller) IndexWithParamTpl(req *ghttp.Request) {
}

func (ctl *Controller) GetAllRssResource(req *ghttp.Request) {

	routerDataList := feedService.GetAllDefinedRouters(gctx.New())
	response.JsonExit(req, 0, "success", routerDataList)
}

func (ctl *Controller) GetAllFeedChannelInfoList(req *ghttp.Request) {

	var feedChannelInfoList []dao.RSSFeed
	var ctx context.Context = context.TODO()

	feedChannelInfoList = feedService.GetAllChannelInfoList(ctx)

	response.JsonExit(req, 0, "success", feedChannelInfoList)
}
