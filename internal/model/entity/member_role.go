// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberRole is the golang structure for table member_role.
type MemberRole struct {
	MemberId  int64       `json:"memberId"  orm:"member_id"  ` // 用户ID
	RoleId    int64       `json:"roleId"    orm:"role_id"    ` // 角色ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
}
