package engadget

import (
	"context"
	"rsshub/app/component"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetIndexRSS(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	apiUrl := "https://www.engadget.com/rss.xml"

	if resp := component.GetContent(ctx,apiUrl); resp != ""{
		req.Response.WriteXmlExit(resp)
	}
}
