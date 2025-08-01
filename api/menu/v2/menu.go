package v2

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 菜单模型
type MenuModel struct {
	Id          int64       `json:"id"         dc:"菜单ID"`
	Type        int         `json:"type"       dc:"菜单类型: 0目录 1菜单 2外链"`
	Name        string      `json:"name"       dc:"菜单名称"`
	Path        string      `json:"path"       dc:"菜单路径"`
	Title       string      `json:"title"      dc:"菜单标题"`
	Icon        string      `json:"icon,omitempty"       dc:"菜单图标"`
	Sort        int         `json:"sort"       dc:"排序"`
	ShowParent  bool        `json:"showParent" dc:"是否显示父菜单"`
	ShowLink    bool        `json:"showLink"   dc:"是否显示该菜单"`
	KeepAlive   bool        `json:"keepAlive"  dc:"是否缓存"`
	ParentId    *int64      `json:"parentId,omitempty"   dc:"父级菜单ID"`
	ParentTitle string      `json:"parentTitle,omitempty" dc:"父级菜单标题"`
	Redirect    string      `json:"redirect,omitempty"   dc:"重定向"`
	Component   string      `json:"component,omitempty"  dc:"组件路径"`
	FrameSrc    string      `json:"frameSrc,omitempty"   dc:"内嵌地址"`
	Url         string      `json:"url,omitempty"        dc:"外部链接"`
	Status      int         `json:"status"     dc:"状态: 0禁用 1启用"`
	CreatedAt   *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"  dc:"更新时间"`
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
	g.Meta     `path:"/menu/create" method:"post" tags:"菜单管理" summary:"创建菜单"`
	Type       int    `json:"type"       v:"required" dc:"菜单类型: 0目录 1菜单 2外链"`
	Name       string `json:"name"       v:"required" dc:"菜单名称"`
	Path       string `json:"path"       v:"required" dc:"菜单路径"`
	Title      string `json:"title"      v:"required" dc:"菜单标题"`
	Icon       string `json:"icon"       dc:"菜单图标"`
	Sort       int    `json:"sort"       v:"required" dc:"排序"`
	ShowParent bool   `json:"showParent" v:"required" dc:"是否显示父菜单"`
	ShowLink   bool   `json:"showLink"   v:"required" dc:"是否显示该菜单"`
	KeepAlive  bool   `json:"keepAlive"  v:"required" dc:"是否缓存"`
	ParentId   int64  `json:"parentId"   dc:"父级ID"`
	Redirect   string `json:"redirect"   dc:"重定向"`
	Component  string `json:"component"  dc:"组件路径"`
	FrameSrc   string `json:"frameSrc"   dc:"内嵌地址"`
	Url        string `json:"url"        dc:"外部链接"`
	Status     int    `json:"status"     v:"required" dc:"状态: 0禁用 1启用"`
}

type CreateRes struct {
}

// UpdateReq 更新菜单
type UpdateReq struct {
	g.Meta     `path:"/menu/update" method:"put" tags:"菜单管理" summary:"更新菜单"`
	Id         int64  `json:"id"         v:"required" dc:"菜单ID"`
	Type       int    `json:"type"       v:"required" dc:"菜单类型: 0目录 1菜单 2外链"`
	Name       string `json:"name"       v:"required" dc:"菜单名称"`
	Path       string `json:"path"       v:"required" dc:"菜单路径"`
	Title      string `json:"title"      v:"required" dc:"菜单标题"`
	Icon       string `json:"icon"       dc:"菜单图标"`
	Sort       int    `json:"sort"       v:"required" dc:"排序"`
	ShowParent bool   `json:"showParent" v:"required" dc:"是否显示父菜单: 0是 1否"`
	ShowLink   bool   `json:"showLink"   v:"required" dc:"是否显示该菜单: 0是 1否"`
	KeepAlive  bool   `json:"keepAlive"  v:"required" dc:"是否缓存: 0是 1否"`
	ParentId   int64  `json:"parentId"   dc:"父级ID"`
	Redirect   string `json:"redirect"   dc:"重定向"`
	Component  string `json:"component"  dc:"组件路径"`
	FrameSrc   string `json:"frameSrc"   dc:"内嵌地址"`
	Url        string `json:"url"        dc:"外部链接"`
	Status     int    `json:"status"     v:"required" dc:"状态: 0禁用 1启用"`
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
