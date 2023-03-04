package service

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

func TestGetHttpClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Init http client",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotClient := GetHttpClient(); gotClient == nil {
				t.Error("GetHttpClient() = nil")
			}
		})
	}
}

func TestGetContent(t *testing.T) {
	type args struct {
		ctx  context.Context
		link string
	}
	tests := []struct {
		name     string
		args     args
		wantResp string
	}{
		{
			name: "Get http response content from baidu",
			args: args{
				ctx:  gctx.New(),
				link: "https://www.baidu.com/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Get http response content from baidu" {
				if gotResp := GetContent(tt.args.ctx, tt.args.link); gotResp == "" {
					t.Error("Get content from baidu is empty")
				}
			} else if gotResp := GetContent(tt.args.ctx, tt.args.link); gotResp != tt.wantResp {
				t.Errorf("GetContent() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestGetContentByMobile(t *testing.T) {
	type args struct {
		ctx  context.Context
		link string
	}
	tests := []struct {
		name     string
		args     args
		wantResp string
	}{
		{
			name: "Get http response content from baidu",
			args: args{
				ctx:  gctx.New(),
				link: "https://www.baidu.com/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Get http response content from baidu" {
				if gotResp := GetContent(tt.args.ctx, tt.args.link); gotResp == "" {
					t.Error("Get content from baidu is empty")
				}
			} else if gotResp := GetContent(tt.args.ctx, tt.args.link); gotResp != tt.wantResp {
				t.Errorf("GetContent() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
