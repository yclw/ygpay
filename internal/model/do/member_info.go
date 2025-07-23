// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberInfo is the golang structure of table t_member_info for DAO operations like Where/Data.
type MemberInfo struct {
	g.Meta       `orm:"table:t_member_info, do:true"`
	Id           interface{} // 用户ID
	Uid          interface{} // 用户UID
	Username     interface{} // 帐号
	PasswordHash interface{} // 密码哈希
	Avatar       interface{} // 头像
	Sex          interface{} // 性别: 1男 2女 3未知
	Email        interface{} // 邮箱
	Mobile       interface{} // 手机号码
	Address      interface{} // 联系地址
	LastActiveAt *gtime.Time // 最后活跃时间
	Remark       interface{} // 备注
	Sort         interface{} // 排序
	Status       interface{} // 状态: 0禁用 1启用
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
