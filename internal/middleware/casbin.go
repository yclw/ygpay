package middleware

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Casbin 鉴权中间件
func (m *Middleware) Casbin(r *ghttp.Request) {
	ctx := r.Context()
	path := r.URL.Path

	// 不需要验证的路由地址
	if isExceptCasbin(ctx, path) {
		r.Middleware.Next()
		return
	}
}

// 不需要验证的路由地址
func isExceptCasbin(ctx context.Context, path string) bool {
	return false
}
