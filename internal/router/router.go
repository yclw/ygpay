package router

import (
	"context"
	"yclw/ygpay/internal/controller/api"
	"yclw/ygpay/internal/controller/login"
	"yclw/ygpay/internal/controller/member"
	"yclw/ygpay/internal/controller/menu"
	"yclw/ygpay/internal/controller/role"
	"yclw/ygpay/internal/controller/user"
	"yclw/ygpay/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Router 设置路由
func Router(ctx context.Context, group *ghttp.Server) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.DefaultMiddleware.Ctx)      // 上下文中间件
		group.Middleware(middleware.DefaultMiddleware.Response) // 响应中间件
		// group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Middleware(middleware.DefaultMiddleware.Lang) // 语言中间件
		group.Bind(
			api.NewV1(),
			login.NewV1(),
		)
		group.Middleware(middleware.DefaultMiddleware.Jwt) // jwt认证中间件
		group.Bind(
			user.NewV2(),
			member.NewV1(),
			role.NewV2(),
			menu.NewV2(),
		)
	})
}
