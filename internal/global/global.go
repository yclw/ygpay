package global

import (
	"yclw/ygpay/pkg/event"

	"github.com/gogf/gf/v2/os/gcache"
)

// VersionApp 应用版本
const (
	VersionApp       = "2.17.8"
	EventServerClose = "server.close" // 服务关闭事件
)

var (
	cache       *gcache.Cache
	serverEvent *event.SEvent
)

func Cache() *gcache.Cache {
	return cache
}

func ServerEvent() *event.SEvent {
	return serverEvent
}
