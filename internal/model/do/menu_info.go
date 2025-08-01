// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MenuInfo is the golang structure of table t_menu_info for DAO operations like Where/Data.
type MenuInfo struct {
	g.Meta     `orm:"table:t_menu_info, do:true"`
	Id         interface{} // 菜单ID
	Type       interface{} // 菜单类型: 0目录 1菜单 2外链
	Name       interface{} // 菜单名称
	Path       interface{} // 菜单路径
	Title      interface{} // 菜单标题
	Icon       interface{} // 菜单图标
	Sort       interface{} // 排序
	ShowParent interface{} // 是否显示父菜单: 0是 1否
	ShowLink   interface{} // 是否显示该菜单: 0是 1否
	KeepAlive  interface{} // 是否缓存: 0是 1否
	Redirect   interface{} // 重定向
	Component  interface{} // 组件路径
	FrameSrc   interface{} // 内嵌地址
	Url        interface{} // 外部链接
	Status     interface{} // 状态: 0禁用 1启用
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
