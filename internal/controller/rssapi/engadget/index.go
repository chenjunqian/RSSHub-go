package engadget

import (
	"context"
	"rsshub/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetIndexRSS(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	apiUrl := "https://www.engadget.com/rss.xml"

	if resp := service.GetContent(ctx,apiUrl); resp != ""{
		req.Response.WriteXmlExit(resp)
	}
}
