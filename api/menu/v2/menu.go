package v2

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 菜单模型
type MenuModel struct {
	Id                int64       `json:"id" dc:"菜单ID"`
	Pid               int64       `json:"pid" dc:"父级ID"`
	Name              string      `json:"name" dc:"菜单名称"`
	Path              string      `json:"path" dc:"路径"`
	Icon              string      `json:"icon" dc:"图标"`
	Title             string      `json:"title" dc:"标题"`
	ShowParent        bool        `json:"showParent" dc:"是否显示父级"`
	Component         string      `json:"component" dc:"组件"`
	NoShowingChildren bool        `json:"noShowingChildren" dc:"是否显示子级"`
	Value             string      `json:"value" dc:"值"`
	ShowTooltip       bool        `json:"showTooltip" dc:"是否显示提示"`
	ParentId          int64       `json:"parentId" dc:"父级ID"`
	Redirect          string      `json:"redirect" dc:"重定向"`
	Description       string      `json:"description" dc:"描述"`
	Sort              int         `json:"sort" dc:"排序"`
	Status            int         `json:"status" dc:"状态"`
	CreatedAt         *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt         *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// GetListReq 获取菜单列表
type GetListReq struct {
	g.Meta `path:"/menu/list" method:"get" tags:"菜单管理" summary:"获取菜单列表"`
	Page   int  `json:"page" v:"required|min:1" dc:"页码"`
	Size   int  `json:"size" v:"required|between:1,100" dc:"每页条数"`
	Status *int `json:"status" dc:"状态筛选（0:禁用 1:启用）"`
}

type GetListRes struct {
	List  []*MenuModel `json:"list" dc:"菜单列表"`
	Total int          `json:"total" dc:"总条数"`
}

// GetOneReq 获取菜单详情
type GetOneReq struct {
	g.Meta `path:"/menu/one" method:"get" tags:"菜单管理" summary:"获取菜单详情"`
	Id     int64 `json:"id" v:"required" dc:"菜单ID"`
}

type GetOneRes struct {
	*MenuModel
}

// CreateReq 创建菜单
type CreateReq struct {
	g.Meta            `path:"/menu/create" method:"post" tags:"菜单管理" summary:"创建菜单"`
	Pid               int64  `json:"pid" v:"required" dc:"父级ID"`
	Name              string `json:"name" v:"required" dc:"菜单名称"`
	Path              string `json:"path" v:"required" dc:"路径"`
	Icon              string `json:"icon" v:"required" dc:"图标"`
	Title             string `json:"title" v:"required" dc:"标题"`
	ShowParent        int    `json:"showParent" v:"required" dc:"是否显示父级"`
	Component         string `json:"component" v:"required" dc:"组件"`
	NoShowingChildren int    `json:"noShowingChildren" v:"required" dc:"是否显示子级"`
	Value             string `json:"value" v:"required" dc:"值"`
	ShowTooltip       int    `json:"showTooltip" v:"required" dc:"是否显示提示"`
	ParentId          int64  `json:"parentId" v:"required" dc:"父级ID"`
	Redirect          string `json:"redirect" v:"required" dc:"重定向"`
	Description       string `json:"description" v:"required" dc:"描述"`
	Sort              int    `json:"sort" v:"required" dc:"排序"`
	Status            int    `json:"status" v:"required" dc:"状态"`
}

type CreateRes struct {
}

// UpdateReq 更新菜单
type UpdateReq struct {
	g.Meta            `path:"/menu/update" method:"put" tags:"菜单管理" summary:"更新菜单"`
	Id                int64  `json:"id" v:"required" dc:"菜单ID"`
	Pid               int64  `json:"pid" v:"required" dc:"父级ID"`
	Name              string `json:"name" v:"required" dc:"菜单名称"`
	Path              string `json:"path" v:"required" dc:"路径"`
	Icon              string `json:"icon" v:"required" dc:"图标"`
	Title             string `json:"title" v:"required" dc:"标题"`
	ShowParent        int    `json:"showParent" v:"required" dc:"是否显示父级"`
	Component         string `json:"component" v:"required" dc:"组件"`
	NoShowingChildren int    `json:"noShowingChildren" v:"required" dc:"是否显示子级"`
	Value             string `json:"value" v:"required" dc:"值"`
	ShowTooltip       int    `json:"showTooltip" v:"required" dc:"是否显示提示"`
	ParentId          int64  `json:"parentId" v:"required" dc:"父级ID"`
	Redirect          string `json:"redirect" v:"required" dc:"重定向"`
	Description       string `json:"description" v:"required" dc:"描述"`
	Sort              int    `json:"sort" v:"required" dc:"排序"`
	Status            int    `json:"status" v:"required" dc:"状态"`
}

type UpdateRes struct {
}

// DeleteReq 删除菜单
type DeleteReq struct {
	g.Meta `path:"/menu/delete" method:"delete" tags:"菜单管理" summary:"删除菜单"`
	Id     int64 `json:"id" v:"required" dc:"菜单ID"`
}

type DeleteRes struct {
}
