package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfoReq 获取用户信息
type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"用户" summary:"获取登录用户信息"`
}

type UserInfoRes struct {
	Id          int64       `json:"id"                 dc:"用户ID"`
	RoleId      int64       `json:"roleId"             dc:"所属角色"`
	Permissions []string    `json:"permissions"        dc:"权限信息"`
	Username    string      `json:"username"           dc:"用户名"`
	Avatar      string      `json:"avatar"             dc:"头像"`
	Sex         int         `json:"sex"                dc:"性别"`
	Email       string      `json:"email"              dc:"邮箱"`
	Mobile      string      `json:"mobile"             dc:"手机号码"`
	Address     string      `json:"address"            dc:"联系地址"`
	CreatedAt   *gtime.Time `json:"createdAt"          dc:"创建时间"`
	LoginCount  int         `json:"loginCount"  dc:"登录次数"`
	LastLoginAt *gtime.Time `json:"lastLoginAt" dc:"最后登录时间"`
	LastLoginIp string      `json:"lastLoginIp" dc:"最后登录IP"`
}
