// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DictData is the golang structure for table dict_data.
type DictData struct {
	Id        int64       `json:"id"        orm:"id"         ` // 字典数据ID
	Label     string      `json:"label"     orm:"label"      ` // 字典标签
	Value     string      `json:"value"     orm:"value"      ` // 字典键值
	ValueType string      `json:"valueType" orm:"value_type" ` // 键值数据类型：string,int,uint,bool,datetime,date
	Type      string      `json:"type"      orm:"type"       ` // 字典类型
	ListClass string      `json:"listClass" orm:"list_class" ` // 表格回显样式
	IsDefault int         `json:"isDefault" orm:"is_default" ` // 是否为系统默认
	Sort      int         `json:"sort"      orm:"sort"       ` // 字典排序
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
	Status    int         `json:"status"    orm:"status"     ` // 状态
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
}
