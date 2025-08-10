// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleApi is the golang structure for table role_api.
type RoleApi struct {
	RoleId    int64       `json:"roleId"    orm:"role_id"    ` // 角色ID
	ApiId     int64       `json:"apiId"     orm:"api_id"     ` // API ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
}
