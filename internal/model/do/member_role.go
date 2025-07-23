// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MemberRole is the golang structure of table t_member_role for DAO operations like Where/Data.
type MemberRole struct {
	g.Meta   `orm:"table:t_member_role, do:true"`
	MemberId interface{} // 管理员ID
	RoleId   interface{} // 角色ID
}
