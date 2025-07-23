// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// MemberRole is the golang structure for table member_role.
type MemberRole struct {
	MemberId int64 `json:"memberId" orm:"member_id" ` // 管理员ID
	RoleId   int64 `json:"roleId"   orm:"role_id"   ` // 角色ID
}
