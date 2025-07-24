package global

import (
	"context"
	"yclw/ygpay/util/token"

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

// GetAccessTokenConfig 获取本地访问token配置
func GetAccessTokenConfig(ctx context.Context) (conf *token.TokenConfig, err error) {
	err = g.Cfg().MustGet(ctx, "token.access").Scan(&conf)
	return
}

// GetRefreshTokenConfig 获取本地刷新token配置
func GetRefreshTokenConfig(ctx context.Context) (conf *token.TokenConfig, err error) {
	err = g.Cfg().MustGet(ctx, "token.refresh").Scan(&conf)
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
