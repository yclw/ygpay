package middleware

import (
	"context"
	"yclw/ygpay/pkg/contexts"
	"yclw/ygpay/pkg/token"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Jwt 鉴权中间件
func (m *Middleware) Jwt(r *ghttp.Request) {
	ctx := r.Context()
	path := r.URL.Path

	// 不需要验证的路由地址
	if isExceptJwt(ctx, path) {
		r.Middleware.Next()
		return
	}

	// 验证token
	tk := r.GetHeader("Authorization")
	// 去掉Bearer前缀
	if len(tk) > 7 && tk[:7] == "Bearer " {
		tk = tk[7:]
	}
	ok, identity, err := token.AccessJwt.VerifyToken(ctx, tk)
	if err != nil || !ok {
		r.SetError(gerror.Wrap(err, "请先重新登录"))
		return
	}

	// 将用户信息传递到上下文中
	userInfo := contexts.Identity{
		Uid:    identity.Uid,
		RoleId: identity.RoleId,
	}
	contexts.SetUser(r.Context(), &userInfo)

	r.Middleware.Next()
}

func isExceptJwt(ctx context.Context, path string) bool {
	except := []string{}

	for _, v := range except {
		if v == path {
			g.Log().Info(ctx, "跳过jwt认证:", v)
			return true
		}
	}

	return false
}
