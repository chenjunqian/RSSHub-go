package main

import (
	_ "rsshub/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"rsshub/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
