package cronJob

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"rsshub/app/component"
	"strings"
)

func RegisterJob() {
	var isProducer bool
	isProducer = g.Cfg().GetBool("guoshao.producer")
	if isProducer {
		_, _ = gcron.AddSingleton("0 */10 * * * *", feedRefreshCronJob)
	}
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
