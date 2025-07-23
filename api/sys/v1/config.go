package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GetConfigReq 获取配置
type GetConfigReq struct {
	g.Meta `path:"/config/get" method:"get" tags:"配置" summary:"获取配置"`
}

type GetConfigRes struct {
	List g.Map `json:"list" dc:"配置列表"`
}

// GetGroupConfigReq 获取指定分组的配置
type GetGroupConfigReq struct {
	g.Meta `path:"/config/get" method:"get" tags:"配置" summary:"获取指定分组的配置"`
	Group  string `json:"group" dc:"分组"`
}

type GetGroupConfigRes struct {
	List g.Map `json:"list" dc:"配置列表"`
}

// UpdateGroupConfigReq 更新指定分组的配置
type UpdateGroupConfigReq struct {
	g.Meta `path:"/config/update" method:"post" tags:"配置" summary:"获取指定分组的配置"`
	Group  string `json:"group"`
	List   g.Map  `json:"list"`
}

type UpdateGroupConfigRes struct {
}

// SiteConfigReq 获取配置
type SiteConfigReq struct {
	g.Meta `path:"/config/site" method:"get" tags:"后台基础" summary:"获取配置"`
}

type SiteConfigRes struct {
	Version string `json:"version"        dc:"系统版本"`
	WsAddr  string `json:"wsAddr"         dc:"客户端websocket地址"`
	Domain  string `json:"domain"         dc:"对外域名"`
	Mode    string `json:"mode"           dc:"运行模式"`
}
