package service

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
)

func GetHttpClient() (client *gclient.Client) {

	client = g.Client()
	client.SetTimeout(time.Second * 60)

	return
}

func GetContent(ctx context.Context, link string) (resp string) {
	var (
		client *gclient.Client
	)
	client = GetHttpClient()
	r, err := client.SetHeaderMap(getHeaders()).Get(ctx, link)
	defer r.Close()
	if err != nil {
		g.Log().Line().Error(ctx, err)
		return
	}

	resp = r.ReadAllString()
	return
}

func GetContentByMobile(ctx context.Context, link string) (resp string) {
	var (
		client *gclient.Client
	)
	client = GetHttpClient()
	resp = client.SetHeaderMap(getMobileHeader()).GetContent(ctx, link)

	return
}

func PostContentByMobile(ctx context.Context, link string, data ...interface{}) (resp string) {
	var (
		client *gclient.Client
	)
	client = GetHttpClient()
	resp = client.SetHeaderMap(getMobileHeader()).PostContent(ctx, link, data)

	return
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func getMobileHeader() (headers map[string]string) {
	headers = make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1"
	return
}
