// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleMenu is the golang structure for table role_menu.
type RoleMenu struct {
	RoleId    int64       `json:"roleId"    orm:"role_id"    ` // 角色ID
	MenuId    int64       `json:"menuId"    orm:"menu_id"    ` // 菜单ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
}
