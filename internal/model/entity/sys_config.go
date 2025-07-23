// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure for table sys_config.
type SysConfig struct {
	Id          int64       `json:"id"          orm:"id"          ` // 配置ID
	Group       string      `json:"group"       orm:"group"       ` // 配置分组
	Key         string      `json:"key"         orm:"key"         ` // 参数键名
	Value       string      `json:"value"       orm:"value"       ` // 参数值
	Description string      `json:"description" orm:"description" ` // 配置描述
	Sort        int         `json:"sort"        orm:"sort"        ` // 排序
	Status      int         `json:"status"      orm:"status"      ` // 状态: 0禁用 1启用
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  ` // 更新时间
}
