package engadget

import (
	"rsshub/app/component"

	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetIndexRSS(req *ghttp.Request) {
	apiUrl := "https://www.engadget.com/rss.xml"

	if resp := component.GetContent(apiUrl); resp != ""{
		_ = req.Response.WriteXmlExit(resp)
	}
}
