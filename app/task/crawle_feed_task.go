package task

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gorilla/feeds"
	feedService "rsshub/app/service/feed"
)

func CallRSSApi(address, route string) (err error) {

	apiUrl := "http://localhost" + address + route
	if _, err = g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err != nil {
		glog.Line().Println("Feed refresh cron job error : ", err)
	}

	return err
}

func StoreFeed(feed, tag, rsshubLink string) (err error) {
	var (
		feedObj  *feeds.Feed
		tagArray []string
	)
	feedObj = new(feeds.Feed)
	tagArray = make([]string, 0)
	_ = gjson.DecodeTo(feed, feedObj)
	_ = gjson.DecodeTo(tag, tagArray)
	err = feedService.AddFeedChannelAndItem(feedObj, tagArray, rsshubLink)
	if err != nil {
		glog.Line().Println(err)
	}
	return err
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}
