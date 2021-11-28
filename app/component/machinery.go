package component

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/backends/result"
	machineryConfig "github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gogf/gf/frame/g"
)

var machineryServer *machinery.Server
var machineryWorker *machinery.Worker

func InitMachinery() {
	g.Log().Info("init machinery ")

	host := g.Cfg().GetString("machinery.redis.host")
	port := g.Cfg().GetString("machinery.redis.port")
	password := g.Cfg().GetString("machinery.redis.pass")

	var cnf = &machineryConfig.Config{
		Broker:        "redis://" + password + "@" + host + ":" + port + "/1",
		DefaultQueue:  "machinery_tasks",
		ResultBackend: "redis://" + password + "@" + host + ":" + port + "/2",
	}
	var serverError error
	machineryServer, serverError = machinery.NewServer(cnf)
	if serverError != nil {
		g.Log().Error("init machinery server failed : ", serverError)
	}

	go func() {
		var err error
		machineryWorker = machineryServer.NewWorker("rsshub_work", 0)
		err = machineryWorker.Launch()
		if err != nil {
			g.Log().Error("init machinery worker failed : ", err)
		}
	}()
}

func SendCallRSSApiTask(address, route string) {
	var signature = &tasks.Signature{
		Name: "CallRSSApi",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: address,
			},
			{
				Type:  "string",
				Value: route,
			},
		},
	}

	var (
		asyncResult *result.AsyncResult
		err         error
	)
	asyncResult, err = machineryServer.SendTask(signature)
	if err != nil {
		g.Log().Error("failed to send task ", err)
	} else if asyncResult.GetState().IsSuccess() {
		g.Log().Info("send task success with signature : ", signature)
	}
}

func SendStoreFeedTask(feed, tag, rsshubLink string) {
	var signature = &tasks.Signature{
		Name: "StoreFeed",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: feed,
			},
			{
				Type:  "string",
				Value: tag,
			},
			{
				Type:  "string",
				Value: rsshubLink,
			},
		},
	}

	var (
		asyncResult *result.AsyncResult
		err         error
	)
	asyncResult, err = machineryServer.SendTask(signature)
	if err != nil {
		g.Log().Error("failed to send task ", err)
	} else if asyncResult.GetState().IsSuccess() {
		g.Log().Info("send task success with signature : ", signature)
	}
}

func GetMachineryServer() *machinery.Server {
	return machineryServer
}
