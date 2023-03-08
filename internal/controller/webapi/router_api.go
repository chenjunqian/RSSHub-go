package webapi

import (
	"rsshub/internal/dao"
	"rsshub/internal/service/feed"
	response "rsshub/internal/service/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"golang.org/x/net/context"
)

func (ctl *Controller) IndexTpl(req *ghttp.Request) {

}

func (ctl *Controller) IndexWithParamTpl(req *ghttp.Request) {
}

func (ctl *Controller) GetAllRssResource(req *ghttp.Request) {

	routerDataList := feed.GetAllDefinedRouters(gctx.New())
	response.JsonExit(req, 0, "success", routerDataList)
}

func (ctl *Controller) GetAllFeedChannelInfoList(req *ghttp.Request) {

	var feedChannelInfoList []dao.RSSFeed
	var ctx context.Context = context.TODO()

	feedChannelInfoList = feed.GetAllChannelInfoList(ctx)

	response.JsonExit(req, 0, "success", feedChannelInfoList)
}
