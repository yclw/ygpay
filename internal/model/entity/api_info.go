// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ApiInfo is the golang structure for table api_info.
type ApiInfo struct {
	Id          int64       `json:"id"          orm:"id"          ` // API ID
	Name        string      `json:"name"        orm:"name"        ` // API名称
	Path        string      `json:"path"        orm:"path"        ` // API路径
	Method      string      `json:"method"      orm:"method"      ` // API方法
	GroupName   string      `json:"groupName"   orm:"group_name"  ` // API分组
	Description string      `json:"description" orm:"description" ` // API描述
	NeedAuth    int         `json:"needAuth"    orm:"need_auth"   ` // 是否需要认证: 0否 1是
	RateLimit   int         `json:"rateLimit"   orm:"rate_limit"  ` // 限流次数/分钟
	Sort        int         `json:"sort"        orm:"sort"        ` // 排序
	Status      int         `json:"status"      orm:"status"      ` // 状态: 0禁用 1启用
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  ` // 更新时间
}
