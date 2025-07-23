// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NoticeRead is the golang structure for table notice_read.
type NoticeRead struct {
	Id        int64       `json:"id"        orm:"id"         ` // 记录ID
	NoticeId  int64       `json:"noticeId"  orm:"notice_id"  ` // 公告ID
	MemberId  int64       `json:"memberId"  orm:"member_id"  ` // 会员ID
	Clicks    int         `json:"clicks"    orm:"clicks"     ` // 已读次数
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 阅读时间
}
