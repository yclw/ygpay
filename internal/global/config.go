package global

import (
	"context"
	"yclw/ygpay/pkg/token"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取运行模式
func GetMode(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "system.mode").String()
}

// 获取时区配置
func GetTimeZone(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "system.timezone").String()
}

// GetLanguage 获取语言配置
func GetLanguage(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "system.language").String()
}

// GetLoadToken 获取本地token配置
func GetLoadToken(ctx context.Context) (conf *token.TokenConfig, err error) {
	err = g.Cfg().MustGet(ctx, "token").Scan(&conf)
	return
}

// AppName 应用名称
func AppName(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "system.appName").String()
}

// Debug debug
func Debug(ctx context.Context) bool {
	return g.Cfg().MustGet(ctx, "system.debug").Bool()
}
