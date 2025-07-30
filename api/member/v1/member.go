package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type MemberModel struct {
	Uid          string
	RoleId       int64
	Username     string
	PasswordHash string
	Nickname     string
	Avatar       string
	Sex          int
	Email        string
	Mobile       string
	Address      string
	LastActiveAt *gtime.Time
	Remark       string
	Sort         int
	Status       int
	CreatedAt    *gtime.Time
	UpdatedAt    *gtime.Time
	LoginCount   int
	LastLoginAt  *gtime.Time
	LastLoginIp  string
}

// GetListReq 获取用户列表
type GetListReq struct {
	g.Meta `path:"/member/list" method:"get" tags:"用户管理" summary:"获取用户列表"`
}

type GetListRes struct {
	List []*MemberModel
}

// GetOneReq 获取用户详情
type GetOneReq struct {
	g.Meta `path:"/member/one" method:"get" tags:"用户管理" summary:"获取用户详情"`
}

type GetOneRes struct {
	Member *MemberModel
}

// CreateReq 创建用户
type CreateReq struct {
	g.Meta   `path:"/member/create" method:"post" tags:"用户管理" summary:"创建用户"`
	Username string `json:"username"            dc:"用户名"`
	Nickname string `json:"nickname"            dc:"昵称"`
	Password string `json:"password"            dc:"密码"`
	RoleId   int64  `json:"roleId"              dc:"角色ID"`
	Avatar   string `json:"avatar"    dc:"头像"`
	Sex      int    `json:"sex"       dc:"性别"`
	Email    string `json:"email"     dc:"邮箱"`
	Mobile   string `json:"mobile"    dc:"手机号码"`
	Address  string `json:"address"   dc:"联系地址"`
	Remark   string `json:"remark"    dc:"备注"`
	Sort     int    `json:"sort"      dc:"排序"`
	Status   int    `json:"status"    dc:"状态"`
}

type CreateRes struct {
}

// UpdateReq 更新用户
type UpdateReq struct {
	g.Meta   `path:"/member/update" method:"put" tags:"用户管理" summary:"更新用户"`
	Uid      string `json:"uid" v:"required"                 dc:"用户ID"`
	Username string `json:"username" v:"required"  dc:"用户名"`
	Nickname string `json:"nickname" v:"required"  dc:"昵称"`
	Avatar   string `json:"avatar" v:"required"    dc:"头像"`
	Sex      int    `json:"sex" v:"required"       dc:"性别"`
	Email    string `json:"email" v:"required"     dc:"邮箱"`
	Mobile   string `json:"mobile" v:"required"    dc:"手机号码"`
	Address  string `json:"address" v:"required"   dc:"联系地址"`
	Remark   string `json:"remark" v:"required"    dc:"备注"`
	Sort     int    `json:"sort" v:"required"      dc:"排序"`
	Status   int    `json:"status" v:"required"    dc:"状态"`
	Password string `json:"password" v:"required"  dc:"密码"`
	RoleId   int64  `json:"roleId" v:"required"    dc:"角色ID"`
}

type UpdateRes struct {
}

// DeleteReq 删除用户
type DeleteReq struct {
	g.Meta `path:"/member/delete" method:"delete" tags:"用户管理" summary:"删除用户"`
}

type DeleteRes struct {
}
