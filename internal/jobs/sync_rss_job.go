package jobs

import (
	"context"
	"rsshub/internal/model"
	"rsshub/internal/service"
	feedService "rsshub/internal/service/feed"
	"strings"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mmcdole/gofeed"
)

func RegisterJob() {
	doSync(doNonAsyncRefreshRSSHub)
}

func doSync(f func()) {
	go func() {
		var freshStartTime = time.Now()
		var refreshHoldTime = time.Minute * 40
		for {
			time.Sleep(time.Minute * 1)
			freshStartTime = time.Now()
			f()
			if time.Now().Sub(freshStartTime) < refreshHoldTime {
				time.Sleep(time.Minute * 60)
			}
		}
	}()
}

func doNonAsyncRefreshRSSHub() {
	var (
		err          error
		ctx          context.Context = context.Background()
		routerLength int             = 0
		routers      []model.RouterInfoData
		rsshubHost   *gvar.Var
	)

	rsshubHost, err = g.Cfg().Get(ctx, "rsshub.baseUrl")
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	routers = feedService.GetAllDefinedRouters(ctx)
	if len(routers) > 0 {
		routerLength = len(routers)
		for index, router := range routers {
			g.Log().Infof(ctx, "Processed %d/%d feed refresh\n", index, routerLength)

			if strings.Contains(router.Route, ":") || strings.Contains(router.Route, "api/") {
				continue
			}
			var (
				apiUrl string
				resp   string
				err    error
				feed   *gofeed.Feed
			)
			apiUrl = rsshubHost.String() + router.Route
			if resp = service.GetContent(ctx, apiUrl); resp == "" {
				g.Log().Errorf(ctx, "get content from url failed, api Url is : %s", apiUrl)
				continue
			}
			fp := gofeed.NewParser()
			feed, err = fp.ParseString(resp)
			if err != nil {
				g.Log().Errorf(ctx, "Parse RSS response error : %s;\nfeed resp: %s;\nAPI url : %s\n", gjson.MustEncode(err), resp, apiUrl)
				continue
			}

			if len(feed.Items) == 0 {
				continue
			}

			err = feedService.AddFeedChannelAndItem(ctx, feed, router.Route)
			if err != nil {
				g.Log().Error(ctx, "Add feed channel and item error : ", gjson.MustEncode(err))
				continue
			}

		}
	} else {
		g.Log().Error(ctx, "The rsshub routers is empty. please check with rsshubHost is : ", rsshubHost)
	}
}