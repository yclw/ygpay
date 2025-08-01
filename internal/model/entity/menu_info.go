// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MenuInfo is the golang structure for table menu_info.
type MenuInfo struct {
	Id         int64       `json:"id"         orm:"id"         ` // 菜单ID
	Type       int         `json:"type"       orm:"type"       ` // 菜单类型: 0目录 1菜单 2外链
	Name       string      `json:"name"       orm:"name"       ` // 菜单名称
	Path       string      `json:"path"       orm:"path"       ` // 菜单路径
	Title      string      `json:"title"      orm:"title"      ` // 菜单标题
	Icon       string      `json:"icon"       orm:"icon"       ` // 菜单图标
	Sort       int         `json:"sort"       orm:"sort"       ` // 排序
	ShowParent int         `json:"showParent" orm:"showParent" ` // 是否显示父菜单: 0是 1否
	ShowLink   int         `json:"showLink"   orm:"showLink"   ` // 是否显示该菜单: 0是 1否
	KeepAlive  int         `json:"keepAlive"  orm:"keepAlive"  ` // 是否缓存: 0是 1否
	Redirect   string      `json:"redirect"   orm:"redirect"   ` // 重定向
	Component  string      `json:"component"  orm:"component"  ` // 组件路径
	FrameSrc   string      `json:"frameSrc"   orm:"frameSrc"   ` // 内嵌地址
	Url        string      `json:"url"        orm:"url"        ` // 外部链接
	Status     int         `json:"status"     orm:"status"     ` // 状态: 0禁用 1启用
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at" ` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at" ` // 更新时间
}
