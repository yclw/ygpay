// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberInfo is the golang structure for table member_info.
type MemberInfo struct {
	Id           int64       `json:"id"           orm:"id"             ` // 用户ID
	Uid          string      `json:"uid"          orm:"uid"            ` // 用户UID
	Username     string      `json:"username"     orm:"username"       ` // 帐号
	PasswordHash string      `json:"passwordHash" orm:"password_hash"  ` // 密码哈希
	Avatar       string      `json:"avatar"       orm:"avatar"         ` // 头像
	Sex          int         `json:"sex"          orm:"sex"            ` // 性别: 1男 2女 3未知
	Email        string      `json:"email"        orm:"email"          ` // 邮箱
	Mobile       string      `json:"mobile"       orm:"mobile"         ` // 手机号码
	Address      string      `json:"address"      orm:"address"        ` // 联系地址
	LastActiveAt *gtime.Time `json:"lastActiveAt" orm:"last_active_at" ` // 最后活跃时间
	Remark       string      `json:"remark"       orm:"remark"         ` // 备注
	Sort         int         `json:"sort"         orm:"sort"           ` // 排序
	Status       int         `json:"status"       orm:"status"         ` // 状态: 0禁用 1启用
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"     ` // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"     ` // 更新时间
	Nickname     string      `json:"nickname"     orm:"nickname"       ` // 昵称
}
