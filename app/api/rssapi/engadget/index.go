package engadget

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetIndexRSS(req *ghttp.Request) {
	apiUrl := "https://www.engadget.com/rss.xml"

	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		_ = req.Response.WriteXmlExit(resp.ReadAllString())
	}
}
