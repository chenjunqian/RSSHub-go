package main

import (
	"rsshub/app/cronJob"
	_ "rsshub/boot"
	_ "rsshub/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	cronJob.RegisterJob()
	g.Server().Run()
}
