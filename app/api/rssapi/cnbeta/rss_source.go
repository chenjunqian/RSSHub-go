package cnbeta

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetRSSSource(req *ghttp.Request) {
	apiUrl := "https://rss.cnbeta.com/"

	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		_ = req.Response.WriteXmlExit(resp.ReadAllString())
	}
}
