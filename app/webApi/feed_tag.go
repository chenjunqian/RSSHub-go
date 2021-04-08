package webApi

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"rsshub/app/service"
	response "rsshub/middleware"
)

func (ctl *Controller) GetFeedTag(req *ghttp.Request) {
	var reqData *FeedTagReqData
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
	tagList := service.GetFeedTag(reqData.Start, reqData.Size)
	response.JsonExit(req, 0, "success", tagList)
}
