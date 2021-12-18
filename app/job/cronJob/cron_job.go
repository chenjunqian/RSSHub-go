package cronJob

import (
	"rsshub/app/component"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func RegisterJob() {
	
	if !g.Cfg().GetBool("guoshao.autoRefreshFeed") {
		return
	}
	
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
		// wait for GoFrame router init
		time.Sleep(time.Second * 5)
		for {
			freshStartTime = time.Now()
			doNonAsyncRefreshFeed()
			if time.Now().Sub(freshStartTime) < refreshHoldTime {
				time.Sleep(time.Minute * 60)
			}
		}
	}()
}

func doNonAsyncRefreshFeed() {
	var (
		routerArray  []ghttp.RouterItem
		routerLength int
	)
	routerArray = g.Server().GetRouterArray()
	if len(routerArray) > 0 {
		routerLength = len(routerArray)
		for index, router := range routerArray {
			if strings.Contains(router.Route, ":") || strings.Contains(router.Route, "rss/api/") {
				continue
			}
			var (
				apiUrl string
				err    error
				resp   *ghttp.ClientResponse
			)
			apiUrl = "http://localhost" + router.Address + router.Route
			if resp, err = component.GetHttpClient().SetHeaderMap(getHeaders()).Get(apiUrl); err != nil {
				g.Log().Error("Feed refresh cron job error : ", err)
			}
			if resp != nil {
				_ = resp.Close()
			}
			g.Log().Infof("Processed %d/%d feed refresh\n", index, routerLength)
		}
	}
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}
