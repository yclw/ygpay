// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleInfo is the golang structure for table role_info.
type RoleInfo struct {
	Id        int64       `json:"id"        orm:"id"         ` // 角色ID
	Name      string      `json:"name"      orm:"name"       ` // 角色名称
	Key       string      `json:"key"       orm:"key"        ` // 角色权限字符串
	Pid       int64       `json:"pid"       orm:"pid"        ` // 上级角色ID
	Level     int         `json:"level"     orm:"level"      ` // 关系树等级
	Tree      string      `json:"tree"      orm:"tree"       ` // 关系树
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
	Sort      int         `json:"sort"      orm:"sort"       ` // 排序
	Status    int         `json:"status"    orm:"status"     ` // 角色状态
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
}
