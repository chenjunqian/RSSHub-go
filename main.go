package main

import (
	_ "rsshub/boot"
	_ "rsshub/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
