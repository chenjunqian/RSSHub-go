package cronJob

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"rsshub/app/component"
	"strings"
	"time"
)

func RegisterJob() {
	if g.Cfg().GetBool("guoshao.asyncRefreshFeed") {
		if g.Cfg().GetBool("guoshao.producer") {
			asyncRefreshFeed()
		}
	} else {
		nonAsyncRefreshFeed()
	}
}

func asyncRefreshFeed() {
	go func() {
		var freshStartTime = time.Now()
		var refreshHoldTime = time.Minute * 40
		for {
			time.Sleep(time.Minute)
			if component.IsAllMachineryTaskDone() {
				if time.Now().Sub(freshStartTime) < refreshHoldTime {
					time.Sleep(time.Minute * 60)
				}
				freshStartTime = time.Now()
				doAsyncRefreshFeed()
			}
		}
	}()
}

func doAsyncRefreshFeed() {
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

func nonAsyncRefreshFeed() {
	go func() {
		var freshStartTime = time.Now()
		var refreshHoldTime = time.Minute * 40
		for {
			if time.Now().Sub(freshStartTime) < refreshHoldTime {
				time.Sleep(time.Minute * 60)
			}
			freshStartTime = time.Now()
			doNonAsyncRefreshFeed()
		}
	}()
}

func doNonAsyncRefreshFeed() {
	routerArray := g.Server().GetRouterArray()
	if len(routerArray) > 0 {
		for _, router := range routerArray {
			if strings.Contains(router.Route, ":") || strings.Contains(router.Route, "rss/api/") {
				continue
			}
			var (
				apiUrl string
				err    error
			)
			apiUrl = "http://localhost" + router.Address + router.Route
			if _, err = component.GetHttpClient().SetHeaderMap(getHeaders()).Get(apiUrl); err != nil {
				glog.Line().Println("Feed refresh cron job error : ", err)
			}

		}
	}
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}
