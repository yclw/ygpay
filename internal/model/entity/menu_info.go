// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MenuInfo is the golang structure for table menu_info.
type MenuInfo struct {
	Id             int64       `json:"id"             orm:"id"              ` // 菜单ID
	Pid            int64       `json:"pid"            orm:"pid"             ` // 父菜单ID
	Level          int         `json:"level"          orm:"level"           ` // 关系树等级
	Tree           string      `json:"tree"           orm:"tree"            ` // 关系树
	Title          string      `json:"title"          orm:"title"           ` // 菜单名称
	Name           string      `json:"name"           orm:"name"            ` // 名称编码
	Path           string      `json:"path"           orm:"path"            ` // 路由地址
	Icon           string      `json:"icon"           orm:"icon"            ` // 菜单图标
	Type           int         `json:"type"           orm:"type"            ` // 菜单类型（1目录 2菜单 3按钮）
	Redirect       string      `json:"redirect"       orm:"redirect"        ` // 重定向地址
	Permissions    string      `json:"permissions"    orm:"permissions"     ` // 菜单包含权限集合
	PermissionName string      `json:"permissionName" orm:"permission_name" ` // 权限名称
	Component      string      `json:"component"      orm:"component"       ` // 组件路径
	AlwaysShow     int         `json:"alwaysShow"     orm:"always_show"     ` // 取消自动计算根路由模式
	ActiveMenu     string      `json:"activeMenu"     orm:"active_menu"     ` // 高亮菜单编码
	IsRoot         int         `json:"isRoot"         orm:"is_root"         ` // 是否跟路由
	IsFrame        int         `json:"isFrame"        orm:"is_frame"        ` // 是否内嵌
	FrameSrc       string      `json:"frameSrc"       orm:"frame_src"       ` // 内联外部地址
	KeepAlive      int         `json:"keepAlive"      orm:"keep_alive"      ` // 缓存该路由
	Hidden         int         `json:"hidden"         orm:"hidden"          ` // 是否隐藏
	Affix          int         `json:"affix"          orm:"affix"           ` // 是否固定
	Sort           int         `json:"sort"           orm:"sort"            ` // 排序
	Remark         string      `json:"remark"         orm:"remark"          ` // 备注
	Status         int         `json:"status"         orm:"status"          ` // 菜单状态
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      ` // 更新时间
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      ` // 创建时间
}
