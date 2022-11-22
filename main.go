package main

import (
	_ "rsshub/boot"
	_ "rsshub/router"

	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	g.Server().Run()
}
