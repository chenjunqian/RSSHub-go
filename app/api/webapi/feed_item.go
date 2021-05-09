package webapi

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"rsshub/app/service"
	response "rsshub/middleware"
)

func (ctl *Controller) GetFeedItemByChannelId(req *ghttp.Request) {
	var reqData *FeedItemListByChannelIdReqData
	if err := req.Parse(&reqData); err != nil {
		if v, ok := err.(*gvalid.Error); ok {
			response.JsonExit(req, 1, v.FirstString())
		} else {
			response.JsonExit(req, 1, err.Error())
		}
	}
	if reqData.Size == 0 {
		reqData.Size = 10
	}

	tagList := service.GetFeedItemByChannelId(reqData.Start, reqData.Size, reqData.ChannelId)
	response.JsonExit(req, 0, "success", tagList)
}

func (ctl *Controller) GetFeedItemListByUserId(req *ghttp.Request) {
	var reqData *FeedItemListByUserIdReqData
	if err := req.Parse(&reqData); err != nil {
		if v, ok := err.(*gvalid.Error); ok {
			response.JsonExit(req, 1, v.FirstString())
		} else {
			response.JsonExit(req, 1, err.Error())
		}
	}
	if reqData.Size == 0 {
		reqData.Size = 10
	}

	itemList := service.GetFeedItemListByUserId(reqData.UserId, reqData.Start, reqData.Size)
	response.JsonExit(req, 0, "success", itemList)
}
