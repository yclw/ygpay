// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MenuInfo is the golang structure for table menu_info.
type MenuInfo struct {
	Id                int64       `json:"id"                orm:"id"                ` // 菜单ID
	Name              string      `json:"name"              orm:"name"              ` // 菜单名称
	Path              string      `json:"path"              orm:"path"              ` // 菜单路径
	Icon              string      `json:"icon"              orm:"icon"              ` // 菜单图标
	Title             string      `json:"title"             orm:"title"             ` // 菜单标题
	ShowParent        int         `json:"showParent"        orm:"showParent"        ` // 是否显示父菜单: 0是 1否
	Component         string      `json:"component"         orm:"component"         ` // 组件路径
	NoShowingChildren int         `json:"noShowingChildren" orm:"noShowingChildren" ` // 是否显示子菜单: 0是 1否
	Value             string      `json:"value"             orm:"value"             ` // 菜单值
	ShowTooltip       int         `json:"showTooltip"       orm:"showTooltip"       ` // 是否显示提示: 0是 1否
	ParentId          int64       `json:"parentId"          orm:"parentId"          ` // 父菜单ID
	Redirect          string      `json:"redirect"          orm:"redirect"          ` // 重定向
	Description       string      `json:"description"       orm:"description"       ` // 菜单描述
	Sort              int         `json:"sort"              orm:"sort"              ` // 排序
	Status            int         `json:"status"            orm:"status"            ` // 状态: 0禁用 1启用
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"        ` // 创建时间
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"        ` // 更新时间
}
