package boot

import (
	_ "rsshub/packed"

	"rsshub/app/component"
	"rsshub/app/job/cronJob"
	"rsshub/app/job/task"

	"github.com/RichardKnop/machinery/v1"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/genv"
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
	component.InitDatabase()
	component.InitES()
	initMachinery()
	cronJob.RegisterJob()
}

func setCookiesToRedis() {
	cookiesMap := g.Cfg().GetMap("cookies")

	for key := range cookiesMap {
		g.Redis().DoVar("SET", key, cookiesMap[key])
	}
}

func initMachinery() {
	if !g.Cfg().GetBool("guoshao.asyncRefreshFeed") {
		return
	}
	component.InitMachinery()

	var isConsumer bool
	isConsumer = g.Cfg().GetBool("guoshao.consumer")
	if !isConsumer {
		return
	}
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

	go func() {
		var err error
		var machineryWorker *machinery.Worker
		machineryWorker = component.GetMachineryServer().NewWorker("rsshub_work", 0)
		err = machineryWorker.Launch()
		if err != nil {
			g.Log().Error("init machinery worker failed : ", err)
		}
	}()
}
