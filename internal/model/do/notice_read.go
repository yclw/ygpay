// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NoticeRead is the golang structure of table t_notice_read for DAO operations like Where/Data.
type NoticeRead struct {
	g.Meta    `orm:"table:t_notice_read, do:true"`
	Id        interface{} // 记录ID
	NoticeId  interface{} // 公告ID
	MemberId  interface{} // 会员ID
	Clicks    interface{} // 已读次数
	UpdatedAt *gtime.Time // 更新时间
	CreatedAt *gtime.Time // 阅读时间
}
