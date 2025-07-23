// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NoticeInfo is the golang structure for table notice_info.
type NoticeInfo struct {
	Id        int64       `json:"id"        orm:"id"         ` // 公告ID
	Title     string      `json:"title"     orm:"title"      ` // 公告标题
	Type      int64       `json:"type"      orm:"type"       ` // 公告类型
	Tag       int         `json:"tag"       orm:"tag"        ` // 标签
	Content   string      `json:"content"   orm:"content"    ` // 公告内容
	Receiver  string      `json:"receiver"  orm:"receiver"   ` // 接收者
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
	Sort      int         `json:"sort"      orm:"sort"       ` // 排序
	Status    int         `json:"status"    orm:"status"     ` // 公告状态
	CreatedBy int64       `json:"createdBy" orm:"created_by" ` // 发送人
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" ` // 修改人
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` // 删除时间
}
