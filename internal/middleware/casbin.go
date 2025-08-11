package middleware

import (
	"context"
	"fmt"

	"yclw/ygpay/internal/global"
	"yclw/ygpay/pkg/contexts"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Casbin 鉴权中间件
func (m *Middleware) Casbin(r *ghttp.Request) {
	ctx := r.Context()
	path := r.URL.Path
	method := r.Method

	// 不需要验证的路由地址
	if isExceptCasbin(ctx, path, method) {
		r.Middleware.Next()
		return
	}

	// 获取用户角色
	roleKey := contexts.GetRoleKey(ctx)

	// casbin 鉴权
	e := global.Casbin()
	ok, err := e.Enforce(roleKey, path, method)
	if err != nil {
		r.SetError(gerror.NewCode(gcode.CodeInternalError, "权限验证系统错误"))
		return
	}
	if !ok {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "权限不足，无法访问该资源"))
		return
	}
	r.Middleware.Next()
}

// 不需要验证的路由地址
func isExceptCasbin(ctx context.Context, path string, method string) bool {
	fmt.Println(path, method)
	return false
}
