package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"rsshub/app/model/biz"
)

func GetFeedTag(start, end int) (tagList []biz.RespFeedTagData) {

	if err := g.DB().Table("rss_feed_tag").
		Fields("name, count(channel_id) as count ").
		Group("name").
		Order("count(channel_id) desc").
		Limit(start, end).
		Structs(&tagList); err != nil {
		glog.Line().Error(err)
	}

	return
}
