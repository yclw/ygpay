package middleware

import (
	"yclw/ygpay/pkg/contexts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (m *Middleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &contexts.Context{
		Session: r.Session,
		Data:    make(g.Map),
	}
	contexts.Init(r, customCtx)
	// 执行下一步请求逻辑
	r.Middleware.Next()
}
