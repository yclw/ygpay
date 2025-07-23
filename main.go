package main

import (
	"yclw/ygpay/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var ctx = gctx.GetInitCtx()
	cmd.Init(ctx)
	cmd.Main.Run(ctx)
}
