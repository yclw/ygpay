package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RoleApiModel struct {
	Id        int64  `json:"id" dc:"APIID"`
	Name      string `json:"name" dc:"API名称"`
	Path      string `json:"path" dc:"API路径"`
	Method    string `json:"method" dc:"API方法"`
	GroupName string `json:"groupName" dc:"API分组"`
	Use       bool   `json:"use" dc:"是否使用"`
}

// GetRoleApiReq 获取角色API
type GetRoleApiReq struct {
	g.Meta `path:"/role/api/get" method:"get" tags:"角色管理" summary:"获取角色API"`
	Id     int64 `json:"id" v:"required" dc:"角色ID"`
}

type GetRoleApiRes struct {
	ApiList []RoleApiModel `json:"apiList" dc:"API列表"`
}

// UpdateRoleApiReq 更新角色API
type UpdateRoleApiReq struct {
	g.Meta  `path:"/role/api/update" method:"post" tags:"角色管理" summary:"更新角色API"`
	Id      int64   `json:"id" v:"required" dc:"角色ID"`
	ApiList []int64 `json:"apiList" v:"required" dc:"API列表"`
}

type UpdateRoleApiRes struct {
}
