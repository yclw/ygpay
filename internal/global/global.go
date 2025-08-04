package global

import (
	"sync"
	"yclw/ygpay/pkg/event"
	"yclw/ygpay/util/i18n"

	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

// VersionApp 应用版本
const (
	VersionApp       = "0.0.1"
	EventServerClose = "server.close" // 服务关闭事件
)

var (
	cache          *gcache.Cache
	serverEvent    *event.SEvent
	language       string
	casbinEnforcer *casbin.Enforcer
	once           sync.Once
)

func Cache() *gcache.Cache {
	once.Do(func() {
		cache = gcache.New()
	})
	return cache
}

func ServerEvent() *event.SEvent {
	return serverEvent
}

func SetLanguage(l string) {
	if !i18n.IsUseLang(l) || l == language {
		return
	}
	language = l
	g.I18n().SetLanguage(language)
}

func SetCasbin(e *casbin.Enforcer) {
	casbinEnforcer = e
}

func Casbin() *casbin.Enforcer {
	return casbinEnforcer
}
