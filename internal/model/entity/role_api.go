// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoleApi is the golang structure for table role_api.
type RoleApi struct {
	RoleId int64 `json:"roleId" orm:"role_id" ` // 角色ID
	ApiId  int64 `json:"apiId"  orm:"api_id"  ` // API ID
}
