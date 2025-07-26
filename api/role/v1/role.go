package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type RoleModel struct {
	Id        int64
	Name      string
	Key       string
	Remark    string
	Sort      int
	Status    int
	CreatedAt *gtime.Time
	UpdatedAt *gtime.Time
}

// GetListReq 获取用户列表
type GetListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"角色管理" summary:"获取角色列表"`
}

type GetListRes struct {
	List []*RoleModel
}

// GetOneReq 获取用户详情
type GetOneReq struct {
	g.Meta `path:"/role/one" method:"get" tags:"角色管理" summary:"获取角色详情"`
}

type GetOneRes struct {
	Member *RoleModel
}

// CreateReq 创建用户
type CreateReq struct {
	g.Meta `path:"/role/create" method:"post" tags:"角色管理" summary:"创建角色"`
	Name   string `json:"name" v:"required" dc:"角色名称"`
	Key    string `json:"key" v:"required" dc:"角色权限字符串"`
	Remark string `json:"remark" v:"required" dc:"备注"`
	Sort   int    `json:"sort" v:"required" dc:"排序"`
	Status int    `json:"status" v:"required" dc:"状态"`
}

type CreateRes struct {
}

// UpdateReq 更新用户
type UpdateReq struct {
	g.Meta `path:"/role/update" method:"put" tags:"角色管理" summary:"更新角色"`
	Id     int64  `json:"id" v:"required" dc:"角色ID"`
	Name   string `json:"name" v:"required" dc:"角色名称"`
	Key    string `json:"key" v:"required" dc:"角色权限字符串"`
	Remark string `json:"remark" v:"required" dc:"备注"`
	Sort   int    `json:"sort" v:"required" dc:"排序"`
	Status int    `json:"status" v:"required" dc:"状态"`
}

type UpdateRes struct {
}

// DeleteReq 删除用户
type DeleteReq struct {
	g.Meta `path:"/role/delete" method:"delete" tags:"角色管理" summary:"删除角色"`
}

type DeleteRes struct {
}
