// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DictType is the golang structure for table dict_type.
type DictType struct {
	Id        int64       `json:"id"        orm:"id"         ` // 字典类型ID
	Pid       int64       `json:"pid"       orm:"pid"        ` // 父类字典类型ID
	Name      string      `json:"name"      orm:"name"       ` // 字典类型名称
	Type      string      `json:"type"      orm:"type"       ` // 字典类型
	Sort      int         `json:"sort"      orm:"sort"       ` // 排序
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
	Status    int         `json:"status"    orm:"status"     ` // 字典类型状态
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
}
