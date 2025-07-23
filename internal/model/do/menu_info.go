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
	g.Meta      `orm:"table:t_menu_info, do:true"`
	Id          interface{} // 菜单ID
	Name        interface{} // 菜单名称
	Path        interface{} // 菜单路径
	Icon        interface{} // 菜单图标
	Component   interface{} // 组件路径
	MenuType    interface{} // 菜单类型: 1目录 2菜单 3按钮
	IsVisible   interface{} // 是否显示: 0否 1是
	Description interface{} // 菜单描述
	Sort        interface{} // 排序
	Status      interface{} // 状态: 0禁用 1启用
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
