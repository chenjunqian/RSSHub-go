package boot

import (
	_ "rsshub/packed"

	"github.com/gogf/gf/frame/g"
)

func init() {
	//app 相关配置
	s := g.Server()
	//GF相关配置 Web Server配置
	g.Log().Stack(true)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	setCookiesToRedis()
}

func setCookiesToRedis() {
	cookiesMap := g.Cfg().GetMap("cookies")

	for key := range cookiesMap {
		g.Redis().DoVar("SET", key, cookiesMap[key])
	}
}
