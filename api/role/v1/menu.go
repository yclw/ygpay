package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoleMenuModel 角色菜单模型
type RoleMenuModel struct {
	Id   int64  `json:"id" dc:"菜单ID"`
	Name string `json:"name" dc:"菜单名称"`
	Sort int    `json:"sort" dc:"排序"`
	Use  bool   `json:"use" dc:"是否使用"`
}

// 角色树模型
type RoleMenuTreeModel struct {
	Children []*RoleMenuTreeModel `json:"children"` // 子节点列表
	*RoleMenuModel
}

// GetRoleMenuReq 获取角色菜单
type GetRoleMenuReq struct {
	g.Meta `path:"/role/menu/get" method:"get" tags:"角色管理" summary:"获取角色菜单"`
	Id     int64 `json:"id" v:"required" dc:"角色ID"`
}

type GetRoleMenuRes struct {
	Tree []*RoleMenuTreeModel `json:"tree" dc:"菜单树"`
}

// UpdateRoleMenuReq 更新角色菜单
type UpdateRoleMenuReq struct {
	g.Meta   `path:"/role/menu/update" method:"post" tags:"角色管理" summary:"更新角色菜单"`
	Id       int64   `json:"id" v:"required" dc:"角色ID"`
	MenuList []int64 `json:"menuList" v:"required" dc:"菜单列表"`
}

type UpdateRoleMenuRes struct {
}
