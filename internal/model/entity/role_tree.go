// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoleTree is the golang structure for table role_tree.
type RoleTree struct {
	Id  int64 `json:"id"  orm:"id"  ` // 角色ID
	Pid int64 `json:"pid" orm:"pid" ` // 父角色ID
}
