package cronJob

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"rsshub/app/component"
	"strings"
)

func RegisterJob() {
	_, _ = gcron.AddSingleton("0 0 * * * *", feedRefreshCronJob)
}

func feedRefreshCronJob() {
	routerArray := g.Server().GetRouterArray()
	if len(routerArray) > 0 {
		for _, router := range routerArray {
			if strings.Contains(router.Route, ":") || strings.Contains(router.Route, "rss/api/") {
				continue
			}

			component.SendCallRSSApiTask(router.Address, router.Route)
		}
	}
}
