package feed

import (
	"context"
	"reflect"
	"rsshub/internal/model"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/mmcdole/gofeed"
)

func Test_assembleFeedChannlAndItem(t *testing.T) {
	var (
		ctx              = gctx.New()
		err              error
		feedXmlStr       string
		rsshubLink       = "/test/router/link"
		fp               *gofeed.Parser
		feed             *gofeed.Feed
		feedChannelModeL model.RssFeedChannel
		feedItemModeList []model.RssFeedItem
	)

	feedXmlStr = gfile.GetContents("./testdata/feed.xml")
	fp = gofeed.NewParser()
	feed, err = fp.ParseString(feedXmlStr)
	if err != nil {
		t.Fatal("Parse feed xml failed : ", err)
	}

	if feed == nil || len(feed.Items) == 0 {
		t.Fatal("Parse feed xml failed, feed is empty : ", err)
	}

	feedChannelModeL, feedItemModeList = assembleFeedChannlAndItem(ctx, feed, rsshubLink)
	if feedChannelModeL.ChannelDesc == "" || feedChannelModeL.Id == "" || feedChannelModeL.Title == "" {
		t.Fatal("assemble feed channel model failed.")
	}

	if len(feedItemModeList) == 0 || feedItemModeList[0].Title == "" || feedItemModeList[0].InputDate == nil || feedItemModeList[0].Link == "" {
		t.Fatal("assemble feed item list model failed.")
	}
}

func Test_getDescriptionThumbnail(t *testing.T) {
	type args struct {
		htmlStr string
	}
	tests := []struct {
		name          string
		args          args
		wantThumbnail string
	}{
		{
			name: "Get content thumbnail",
			args: args{
				htmlStr: gfile.GetContents("./testdata/feed_item_content.html"),
			},
			wantThumbnail: "https://1-im.guokr.com/IiAbcj0VwBimp50wk5ji8G2TWz_SPbqAYxT7pydfZHHkAwAANwIAAEpQ.jpg?imageView2/1/w/555/h/315",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotThumbnail := getDescriptionThumbnail(tt.args.htmlStr); gotThumbnail != tt.wantThumbnail {
				t.Errorf("getDescriptionThumbnail() = %v, want %v", gotThumbnail, tt.wantThumbnail)
			}
		})
	}
}

func TestGetAllDefinedRouters(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name               string
		args               args
		wantRouterDataList []model.RouterInfoData
	}{
		{
			name: "Get all defined router list",
			args: args{
				ctx: gctx.New(),
			},
			wantRouterDataList: []model.RouterInfoData{
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRouterDataList := GetAllDefinedRouters(tt.args.ctx); !reflect.DeepEqual(gotRouterDataList, tt.wantRouterDataList) {
				t.Errorf("GetAllDefinedRouters() = %v, want %v", gotRouterDataList, tt.wantRouterDataList)
			}
		})
	}
}
