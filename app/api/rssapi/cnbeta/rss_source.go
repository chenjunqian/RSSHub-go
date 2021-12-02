package cnbeta

import (
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
)

func (ctl *Controller) GetRSSSource(req *ghttp.Request) {
	apiUrl := "https://rss.cnbeta.com/"

	if resp, err := component.GetHttpClient().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		_ = req.Response.WriteXmlExit(resp.ReadAllString())
	}
}
