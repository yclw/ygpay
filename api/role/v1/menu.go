package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 菜单模型
type MenuModel struct {
	// Id       int64        `json:"id"         dc:"菜单ID"`
	MenuUid  string       `json:"menuUid"    dc:"菜单唯一标识"`
	ParentId int64        `json:"-"   dc:"父级菜单ID"`
	Title    string       `json:"title"      dc:"菜单标题"`
	Sort     int          `json:"-"       dc:"排序"`
	Use      bool         `json:"use"        dc:"是否使用"`
	Children []*MenuModel `json:"children,omitempty" dc:"子菜单"`
}

// GetRoleMenuReq 获取角色菜单
type GetRoleMenuReq struct {
	g.Meta  `path:"/role/menu/get" method:"get" tags:"角色管理" summary:"获取角色菜单"`
	RoleUid string `json:"roleUid" v:"required" dc:"角色唯一标识"`
}

type GetRoleMenuRes struct {
	Tree []*MenuModel `json:"tree" dc:"菜单树"`
}

// UpdateRoleMenuReq 更新角色菜单
type UpdateRoleMenuReq struct {
	g.Meta   `path:"/role/menu/update" method:"post" tags:"角色管理" summary:"更新角色菜单"`
	RoleUid  string   `json:"roleUid" v:"required" dc:"角色唯一标识"`
	MenuList []string `json:"menuList" v:"required" dc:"菜单列表"`
}

type UpdateRoleMenuRes struct {
}
