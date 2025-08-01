package v2

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfoReq 获取用户信息
type GetUserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"用户" summary:"获取登录用户信息"`
}

type GetUserInfoRes struct {
	Uid       string      `json:"uid"                 dc:"用户UID"`
	Nickname  string      `json:"nickname"           dc:"昵称"`
	RoleName  string      `json:"roleName"           dc:"角色名称"`
	Username  string      `json:"username"           dc:"用户名"`
	Avatar    string      `json:"avatar"             dc:"头像"`
	Sex       int         `json:"sex"                dc:"性别"`
	Email     string      `json:"email"              dc:"邮箱"`
	Mobile    string      `json:"mobile"             dc:"手机号码"`
	Address   string      `json:"address"            dc:"联系地址"`
	CreatedAt *gtime.Time `json:"createdAt"          dc:"创建时间"`
}
