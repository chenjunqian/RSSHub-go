package cnbeta

import (
	"context"
	"rsshub/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetRSSSource(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	apiUrl := "https://rss.cnbeta.com/"

	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		req.Response.WriteXmlExit(resp)
	}
}
