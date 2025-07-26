package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoleMenuModel 角色菜单模型
type RoleMenuModel struct {
	Id   int64
	Name string
	Key  string
	Sort int
	Use  bool
}

// GetRoleMenuReq 获取角色菜单
type GetRoleMenuReq struct {
	g.Meta `path:"/role/menu/get" method:"get" tags:"角色管理" summary:"获取角色菜单"`
	Id     int64 `json:"id" v:"required" dc:"角色ID"`
}

type GetRoleMenuRes struct {
	MenuList []RoleMenuModel `json:"menuList" dc:"菜单列表"`
}

// UpdateRoleMenuReq 更新角色菜单
type UpdateRoleMenuReq struct {
	g.Meta   `path:"/role/menu/update" method:"post" tags:"角色管理" summary:"更新角色菜单"`
	Id       int64   `json:"id" v:"required" dc:"角色ID"`
	MenuList []int64 `json:"menuList" v:"required" dc:"菜单列表"`
}

type UpdateRoleMenuRes struct {
}
