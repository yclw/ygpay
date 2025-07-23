// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleInfo is the golang structure of table t_role_info for DAO operations like Where/Data.
type RoleInfo struct {
	g.Meta    `orm:"table:t_role_info, do:true"`
	Id        interface{} // 角色ID
	Name      interface{} // 角色名称
	Key       interface{} // 角色权限字符串
	Pid       interface{} // 上级角色ID
	Level     interface{} // 关系树等级
	Tree      interface{} // 关系树
	Remark    interface{} // 备注
	Sort      interface{} // 排序
	Status    interface{} // 角色状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
