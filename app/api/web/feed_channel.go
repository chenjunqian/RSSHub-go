package web

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"rsshub/app/service"
	response "rsshub/middleware"
)

func (ctl *Controller) GetFeedChannelByTag(req *ghttp.Request) {
	var reqData *FeedChannelReqData
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
	tagList := service.GetFeedChannelByTag(reqData.Start, reqData.Size, reqData.Name)
	response.JsonExit(req, 0, "success", tagList)
}
