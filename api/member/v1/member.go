package v1

import "github.com/gogf/gf/v2/frame/g"

// GetListReq 获取用户列表
type GetListReq struct {
	g.Meta `path:"/member/list" method:"get" tags:"用户" summary:"获取用户列表"`
}

type GetListRes struct {
}

// GetOneReq 获取用户详情
type GetOneReq struct {
	g.Meta `path:"/member/one" method:"get" tags:"用户" summary:"获取用户详情"`
}

type GetOneRes struct {
}

// CreateReq 创建用户
type CreateReq struct {
	g.Meta       `path:"/member/create" method:"post" tags:"用户" summary:"创建用户"`
	Username     string `json:"username"           dc:"用户名"`
	PasswordHash string `json:"passwordHash"       dc:"密码"`
	RoleId       int64  `json:"roleId"             dc:"角色ID"`
	Avatar       string `json:"avatar"             dc:"头像"`
	Sex          int    `json:"sex"                dc:"性别"`
	Email        string `json:"email"              dc:"邮箱"`
	Mobile       string `json:"mobile"             dc:"手机号码"`
	Address      string `json:"address"            dc:"联系地址"`
	Remark       string `json:"remark"             dc:"备注"`
	Status       int    `json:"status"             dc:"状态"`
}

type CreateRes struct {
}

// UpdateReq 更新用户
type UpdateReq struct {
	g.Meta       `path:"/member/update" method:"put" tags:"用户" summary:"更新用户"`
	Id           int64  `json:"id"                 dc:"用户ID"`
	Username     string `json:"username"           dc:"用户名"`
	PasswordHash string `json:"passwordHash"       dc:"密码"`
	RoleId       int64  `json:"roleId"             dc:"角色ID"`
	Avatar       string `json:"avatar"             dc:"头像"`
	Sex          int    `json:"sex"                dc:"性别"`
	Email        string `json:"email"              dc:"邮箱"`
	Mobile       string `json:"mobile"             dc:"手机号码"`
	Address      string `json:"address"            dc:"联系地址"`
	Remark       string `json:"remark"             dc:"备注"`
	Status       int    `json:"status"             dc:"状态"`
}

type UpdateRes struct {
}

// DeleteReq 删除用户
type DeleteReq struct {
	g.Meta `path:"/member/delete" method:"delete" tags:"用户" summary:"删除用户"`
}

type DeleteRes struct {
}
