package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/genv"
	"rsshub/app/component"
	"rsshub/app/cronJob"
	"rsshub/app/task"
	_ "rsshub/packed"
)

func init() {
	configEvn := genv.Get("GF_GCFG_FILE", "")
	if configEvn != "" {
		g.Cfg().SetFileName(configEvn)
	}
	//app 相关配置
	s := g.Server()
	//GF相关配置 Web Server配置
	g.Log().Stack(true)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	setCookiesToRedis()
	component.InitES()
	component.InitMachinery()
	cronJob.RegisterJob()

	var err error
	err = component.GetMachineryServer().RegisterTask("CallRSSApi", task.CallRSSApi)
	if err != nil {
		g.Log().Error("machinery register task CallRSSApi failed : ", err)
	}
	err = component.GetMachineryServer().RegisterTask("StoreFeed", task.StoreFeed)
	if err != nil {
		g.Log().Error("machinery register task StoreFeed failed : ", err)
	}
	g.Log().Info("machinery register task : ", component.GetMachineryServer().GetRegisteredTaskNames())
}

func setCookiesToRedis() {
	cookiesMap := g.Cfg().GetMap("cookies")

	for key := range cookiesMap {
		g.Redis().DoVar("SET", key, cookiesMap[key])
	}
}
