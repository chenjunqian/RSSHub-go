package engadget

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/component"
)

func (ctl *Controller) GetIndexRSS(req *ghttp.Request) {
	apiUrl := "https://www.engadget.com/rss.xml"

	if resp, err := component.GetHttpClient().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		defer func(resp *ghttp.ClientResponse) {
			err := resp.Close()
			if err != nil {
				g.Log().Error(err)
			}
		}(resp)
		_ = req.Response.WriteXmlExit(resp.ReadAllString())
	}
}
