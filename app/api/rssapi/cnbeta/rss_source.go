package cnbeta

import (
	"context"
	"rsshub/app/component"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetRSSSource(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	apiUrl := "https://rss.cnbeta.com/"

	if resp := component.GetContent(ctx,apiUrl); resp != "" {
		req.Response.WriteXmlExit(resp)
	}
}
