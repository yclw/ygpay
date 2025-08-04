package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// API模型
type ApiModel struct {
	Id     int64  `json:"id" dc:"APIID"`
	Path   string `json:"path" dc:"API路径"`
	Method string `json:"method" dc:"API方法"`
	Sort   int    `json:"-" dc:"排序"`
	Group  string `json:"-" dc:"API分组"`
	Use    bool   `json:"use" dc:"是否使用"`
}

// API分组模型
type ApiGroupModel struct {
	GroupName string      `json:"groupName" dc:"API分组"`
	Children  []*ApiModel `json:"children" dc:"子API"`
}

// GetRoleApiReq 获取角色API
type GetRoleApiReq struct {
	g.Meta `path:"/role/api/get" method:"get" tags:"角色管理" summary:"获取角色API"`
	Id     int64 `json:"id" v:"required" dc:"角色ID"`
}

type GetRoleApiRes struct {
	ApiList []*ApiGroupModel `json:"apiList" dc:"API列表"`
}

// UpdateRoleApiReq 更新角色API
type UpdateRoleApiReq struct {
	g.Meta  `path:"/role/api/update" method:"post" tags:"角色管理" summary:"更新角色API"`
	Id      int64   `json:"id" v:"required" dc:"角色ID"`
	ApiList []int64 `json:"apiList" v:"required" dc:"API列表"`
}

type UpdateRoleApiRes struct {
}
