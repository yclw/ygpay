package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 角色模型
type RoleModel struct {
	RoleUid    string       `json:"roleUid" dc:"角色唯一标识"`
	Name       string       `json:"name" dc:"角色名称"`
	Key        string       `json:"key" dc:"角色权限字符串"`
	ParentId   int64        `json:"-" dc:"父级ID"`
	ParentName string       `json:"parentName" dc:"父级名称"`
	Remark     string       `json:"remark" dc:"备注"`
	Sort       int          `json:"sort" dc:"排序"`
	Status     int          `json:"status" dc:"状态"`
	CreatedAt  *gtime.Time  `json:"createdAt" dc:"创建时间"`
	UpdatedAt  *gtime.Time  `json:"updatedAt" dc:"更新时间"`
	Children   []*RoleModel `json:"children" dc:"子角色"`
}

// GetListReq 获取用户列表
type GetListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"角色管理" summary:"获取角色列表"`
}

type GetListRes struct {
	List []*RoleModel `json:"list" dc:"角色列表"`
	Tree []*RoleModel `json:"tree" dc:"角色树"`
}

// GetOneReq 获取用户详情
type GetOneReq struct {
	g.Meta  `path:"/role/one" method:"get" tags:"角色管理" summary:"获取角色详情"`
	RoleUid string `json:"roleUid" v:"required" dc:"角色唯一标识"`
}

type GetOneRes struct {
	*RoleModel
}

// CreateReq 创建用户
type CreateReq struct {
	g.Meta    `path:"/role/create" method:"post" tags:"角色管理" summary:"创建角色"`
	ParentUid string `json:"parentUid" dc:"父级唯一标识"`
	Name      string `json:"name" v:"required" dc:"角色名称"`
	Key       string `json:"key" v:"required" dc:"角色权限字符串"`
	Remark    string `json:"remark" v:"required" dc:"备注"`
	Sort      int    `json:"sort" v:"required" dc:"排序"`
	Status    int    `json:"status" v:"required" dc:"状态"`
}

type CreateRes struct {
}

// UpdateReq 更新用户
type UpdateReq struct {
	g.Meta    `path:"/role/update" method:"put" tags:"角色管理" summary:"更新角色"`
	RoleUid   string `json:"roleUid" v:"required" dc:"角色唯一标识"`
	ParentUid string `json:"parentUid" dc:"父级唯一标识"`
	Name      string `json:"name" v:"required" dc:"角色名称"`
	Key       string `json:"key" v:"required" dc:"角色权限字符串"`
	Remark    string `json:"remark" v:"required" dc:"备注"`
	Sort      int    `json:"sort" v:"required" dc:"排序"`
	Status    int    `json:"status" v:"required" dc:"状态"`
}

type UpdateRes struct {
}

// DeleteReq 删除用户
type DeleteReq struct {
	g.Meta  `path:"/role/delete" method:"delete" tags:"角色管理" summary:"删除角色"`
	RoleUid string `json:"roleUid" v:"required" dc:"角色唯一标识"`
}

type DeleteRes struct {
}
