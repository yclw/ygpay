// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoleTree is the golang structure of table t_role_tree for DAO operations like Where/Data.
type RoleTree struct {
	g.Meta `orm:"table:t_role_tree, do:true"`
	Id     interface{} // 角色ID
	Pid    interface{} // 父角色ID
}
