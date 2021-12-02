package component

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"time"
)

func GetHttpClient() (client *ghttp.Client) {

	client = g.Client()
	client.SetTimeout(time.Second * 10)

	return
}
