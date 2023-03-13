package cache

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

func TestInitSiteCookies(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Get cookies with no error",
			args: args{
				ctx: gctx.New(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitCache(tt.args.ctx)
			if err := InitSiteCookies(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("InitSiteCookies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetSiteCookies(t *testing.T) {
	type args struct {
		ctx      context.Context
		siteName string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Get zhihu test cookies",
			args: args{
				ctx:      gctx.New(),
				siteName: "zhihu",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitCache(tt.args.ctx)
			InitSiteCookies(tt.args.ctx)
			if got := GetSiteCookies(tt.args.ctx, tt.args.siteName); got == nil {
				t.Error("GetSiteCookies() = nil")
			}
		})
	}
}
