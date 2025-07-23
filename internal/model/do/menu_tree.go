// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MenuTree is the golang structure of table t_menu_tree for DAO operations like Where/Data.
type MenuTree struct {
	g.Meta `orm:"table:t_menu_tree, do:true"`
	Id     interface{} // 菜单ID
	Pid    interface{} // 父菜单ID
}
