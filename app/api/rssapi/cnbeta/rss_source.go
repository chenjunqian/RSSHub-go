package cnbeta

import (
	"rsshub/app/component"

	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetRSSSource(req *ghttp.Request) {
	apiUrl := "https://rss.cnbeta.com/"

	if resp := component.GetContent(apiUrl); resp != "" {
		_ = req.Response.WriteXmlExit(resp)
	}
}
