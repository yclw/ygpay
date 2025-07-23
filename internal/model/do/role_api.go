// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoleApi is the golang structure of table t_role_api for DAO operations like Where/Data.
type RoleApi struct {
	g.Meta `orm:"table:t_role_api, do:true"`
	RoleId interface{} // 角色ID
	ApiId  interface{} // API ID
}
