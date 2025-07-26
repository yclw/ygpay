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
	Group  string `json:"group" v:"required" dc:"分组"`
}

type GetGroupConfigRes struct {
	List g.Map `json:"list" dc:"配置列表"`
}

// UpdateGroupConfigReq 更新指定分组的配置
type UpdateGroupConfigReq struct {
	g.Meta `path:"/config/update" method:"post" tags:"配置" summary:"获取指定分组的配置"`
	Group  string `json:"group" v:"required" dc:"分组"`
	List   g.Map  `json:"list" v:"required" dc:"配置列表"`
}

type UpdateGroupConfigRes struct {
}
