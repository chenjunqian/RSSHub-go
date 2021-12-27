package boot

import (
	_ "rsshub/packed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/genv"
)

func init() {
	configEvn := genv.Get("GF_GCFG_FILE", "")
	if configEvn != "" {
		g.Cfg().SetFileName(configEvn)
	}

	s := g.Server()
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
