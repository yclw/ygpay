// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NoticeInfo is the golang structure of table t_notice_info for DAO operations like Where/Data.
type NoticeInfo struct {
	g.Meta    `orm:"table:t_notice_info, do:true"`
	Id        interface{} // 公告ID
	Title     interface{} // 公告标题
	Type      interface{} // 公告类型
	Tag       interface{} // 标签
	Content   interface{} // 公告内容
	Receiver  interface{} // 接收者
	Remark    interface{} // 备注
	Sort      interface{} // 排序
	Status    interface{} // 公告状态
	CreatedBy interface{} // 发送人
	UpdatedBy interface{} // 修改人
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
