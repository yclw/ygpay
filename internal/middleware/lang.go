package middleware

import (
	"yclw/ygpay/internal/global"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (m *Middleware) Lang(r *ghttp.Request) {
	language := r.GetHeader("X-Language")
	global.SetLanguage(language)
	r.Middleware.Next()
}
