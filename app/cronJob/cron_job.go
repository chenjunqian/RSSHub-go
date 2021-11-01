package cronJob

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/grpool"
	"strings"
)

func RegisterJob() {
	_, _ = gcron.AddSingleton("0 0 * * * *", feedRefreshCronJob)
}

func feedRefreshCronJob() {
	routerArray := g.Server().GetRouterArray()
	pool := grpool.New(2)
	if len(routerArray) > 0 {
		for _, router := range routerArray {
			if strings.Contains(router.Route, ":") || strings.Contains(router.Route, "rss/api/") {
				continue
			}

			tempRout := router
			_ = pool.Add(func() {
				apiUrl := "http://localhost" + tempRout.Address + tempRout.Route
				if _, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err != nil {
					glog.Line().Println("Feed refresh cron job error : ", err)
				}
			})
		}
	}
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}
