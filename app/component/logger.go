package component

import (
	"rsshub/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

var logger *glog.Logger

func InitLogger() {
	logger = glog.New()
	logger.SetConfigWithMap(g.Map{
		"path":  config.GetConfig().Get("log.logPath").String(),
		"level": config.GetConfig().Get("log.level").String(),
	})
}

func GetLogger() *glog.Logger {
	return logger
}
