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
	g.Meta             `orm:"table:t_member_info, do:true"`
	Id                 interface{} // 用户ID
	Username           interface{} // 帐号
	PasswordHash       interface{} // 密码
	Salt               interface{} // 密码盐
	PasswordResetToken interface{} // 密码重置令牌
	RoleId             interface{} // 角色ID
	Avatar             interface{} // 头像
	Sex                interface{} // 性别
	Email              interface{} // 邮箱
	Mobile             interface{} // 手机号码
	Address            interface{} // 联系地址
	LastActiveAt       *gtime.Time // 最后活跃时间
	Remark             interface{} // 备注
	Status             interface{} // 状态
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 修改时间
}
