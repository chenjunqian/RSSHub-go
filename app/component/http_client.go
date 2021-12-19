package component

import (
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func GetHttpClient() (client *ghttp.Client) {

	client = g.Client()
	client.SetTimeout(time.Second * 60)

	return
}

func GetContent(link string) (resp string) {
	var (
		client *ghttp.Client
	)
	client = GetHttpClient()
	resp = client.SetHeaderMap(getHeaders()).GetContent(link)

	return
}

func GetContentByMobile(link string) (resp string) {
	var (
		client *ghttp.Client
	)
	client = GetHttpClient()
	resp = client.SetHeaderMap(getMobileHeader()).GetContent(link)

	return
}

func PostContentByMobile(link string, data ...interface{}) (resp string) {
	var (
		client *ghttp.Client
	)
	client = GetHttpClient()
	resp = client.SetHeaderMap(getMobileHeader()).PostContent(link,data)

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
