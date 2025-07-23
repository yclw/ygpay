package cmd

import (
	"context"
	"yclw/ygpay/internal/router"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			router.Router(ctx, s)
			s.Run()
			return nil
		},
	}
)
