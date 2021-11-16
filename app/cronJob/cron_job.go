package cronJob

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcron"
	"rsshub/app/component"
	"strings"
)

func RegisterJob() {
	_, _ = gcron.AddSingleton("0 */10 * * * *", feedRefreshCronJob)
}

func feedRefreshCronJob() {
	routerArray := g.Server().GetRouterArray()
	if len(routerArray) > 0 {
		for _, router := range routerArray {
			if strings.Contains(router.Route, ":") || strings.Contains(router.Route, "rss/api/") {
				continue
			}

			var (
				tempRout ghttp.RouterItem
			)
			tempRout = router
			component.SendCallRSSApiTask(tempRout.Address, tempRout.Route)
		}
	}
}
