// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// MenuTree is the golang structure for table menu_tree.
type MenuTree struct {
	Id  int64 `json:"id"  orm:"id"  ` // 菜单ID
	Pid int64 `json:"pid" orm:"pid" ` // 父菜单ID
}
