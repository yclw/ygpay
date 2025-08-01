package cmd

import (
	"context"
	"fmt"
	"runtime"
	"yclw/ygpay/internal/global"
	"yclw/ygpay/pkg/token"
	"yclw/ygpay/util/i18n"

	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/golang-jwt/jwt/v5"
)

func Init(ctx context.Context) {
	// 设置gf运行模式
	SetGFMode(ctx)

	// 默认上海时区
	SetTimeZone(ctx)

	// 初始化i18n
	InitI18n(ctx)

	// 输出欢迎信息
	PrintWelcome(ctx)

	// 初始化链路追踪
	InitTrace(ctx)

	// 设置缓存适配器
	InitAdapter(ctx)

	// 初始化token
	InitToken(ctx)
}

// SetGFMode 设置gf运行模式
func SetGFMode(ctx context.Context) {
	mode := global.GetMode(ctx)
	gmode.Set(mode)
}

// SetTimeZone 设置时区
func SetTimeZone(ctx context.Context) {
	zone := global.GetTimeZone(ctx)
	if err := gtime.SetTimeZone(zone); err != nil {
		g.Log().Fatalf(ctx, "时区设置异常 err：%+v", err)
		return
	}
	g.Log().Debug(ctx, "时区设置成功：%v", zone)
}

// SetI18n 设置i18n
func InitI18n(ctx context.Context) {
	language := global.GetLanguageConfig(ctx)

	i18n.AddUseLang(i18n.LangZhCN)
	i18n.AddUseLang(i18n.LangEn)

	if !i18n.IsUseLang(language) {
		language = i18n.LangZhCN
	}

	global.SetLanguage(language)
	g.Log().Debug(ctx, "i18n设置成功：%v", language)
}

// PrintWelcome 输出欢迎信息
func PrintWelcome(ctx context.Context) {
	appName := global.AppName(ctx)
	fmt.Printf("欢迎使用%v！\r\n当前运行环境：%v, 运行根路径为：%v \r\nAPP版本：v%v, gf版本：%v \n", appName, runtime.GOOS, gfile.Pwd(), global.VersionApp, gf.VERSION)
}

// InitTrace 初始化链路追踪
func InitTrace(ctx context.Context) {
	if !g.Cfg().MustGet(ctx, "jaeger.switch").Bool() {
		return
	}

	tp, err := jaeger.Init(global.AppName(ctx), g.Cfg().MustGet(ctx, "jaeger.endpoint").String())
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	global.ServerEvent().Register(global.EventServerClose, func(ctx context.Context, args ...interface{}) {
		_ = tp.Shutdown(ctx)
		g.Log().Debug(ctx, "jaeger closed ..")
	})
}

// InitAdapter 设置缓存适配器
func InitAdapter(ctx context.Context) {
	adapterType := global.CacheAdapterConfig(ctx)
	var adapter gcache.Adapter
	switch adapterType {
	case "redis":
		adapter = gcache.NewAdapterRedis(g.Redis())
	default:
		adapter = gcache.NewAdapterMemory()
	}
	global.Cache().SetAdapter(adapter)
}

// InitToken 初始化token
func InitToken(ctx context.Context) {
	// 初始化访问令牌
	conf, err := global.GetAccessTokenConfig(ctx)
	if err != nil {
		return
	}
	token.AccessJwt = token.NewAccessJwtHandler(conf, jwt.SigningMethodHS256)

	// 初始化刷新令牌
	conf, err = global.GetRefreshTokenConfig(ctx)
	if err != nil {
		return
	}
	token.RefreshJwt = token.NewRefreshJwtHandler(conf, jwt.SigningMethodHS256)
}
