package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type MemberModel struct {
	Uid          string      `json:"uid" dc:"用户ID"`
	RoleId       int64       `json:"roleId" dc:"角色ID"`
	RoleName     string      `json:"roleName,omitempty" dc:"角色名称"`
	Username     string      `json:"username" dc:"用户名"`
	Nickname     string      `json:"nickname" dc:"昵称"`
	Avatar       string      `json:"avatar" dc:"头像"`
	Sex          int         `json:"sex" dc:"性别: 0未知 1男 2女"`
	Email        string      `json:"email" dc:"邮箱"`
	Mobile       string      `json:"mobile" dc:"手机号码"`
	Address      string      `json:"address" dc:"联系地址"`
	LastActiveAt *gtime.Time `json:"lastActiveAt" dc:"最后活跃时间"`
	Remark       string      `json:"remark" dc:"备注"`
	Sort         int         `json:"sort" dc:"排序"`
	Status       int         `json:"status" dc:"状态: 0禁用 1启用"`
	CreatedAt    *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// GetListReq 获取用户列表
type GetListReq struct {
	g.Meta    `path:"/member/list" method:"get" tags:"用户管理" summary:"获取用户列表"`
	Page      int         `json:"page" v:"required|min:1" dc:"页码"`
	Size      int         `json:"size" v:"required|between:1,100" dc:"每页条数"`
	Username  string      `json:"username" dc:"用户名"`
	Nickname  string      `json:"nickname" dc:"昵称"`
	Email     string      `json:"email" dc:"邮箱"`
	Mobile    string      `json:"mobile" dc:"手机号码"`
	RoleId    *int64      `json:"roleId" dc:"角色ID筛选"`
	Sex       *int        `json:"sex" dc:"性别筛选（0:未知 1:男 2:女）"`
	Status    *int        `json:"status" dc:"状态筛选（0:禁用 1:启用）"`
	StartDate *gtime.Time `json:"startDate" dc:"开始日期"`
	EndDate   *gtime.Time `json:"endDate" dc:"结束日期"`
	SortField string      `json:"sortField" dc:"排序字段"`
	SortDesc  bool        `json:"sortDesc" dc:"是否降序"`
}

type GetListRes struct {
	List  []*MemberModel `json:"list" dc:"用户列表"`
	Total int            `json:"total" dc:"总条数"`
}

// GetOneReq 获取用户详情
type GetOneReq struct {
	g.Meta `path:"/member/one" method:"get" tags:"用户管理" summary:"获取用户详情"`
	Uid    string `json:"uid" v:"required" dc:"用户ID"`
}

type GetOneRes struct {
	*MemberModel
}

// CreateReq 创建用户
type CreateReq struct {
	g.Meta   `path:"/member/create" method:"post" tags:"用户管理" summary:"创建用户"`
	Username string `json:"username" v:"required" dc:"用户名"`
	Nickname string `json:"nickname" v:"required" dc:"昵称"`
	Password string `json:"password" v:"required" dc:"密码"`
	RoleId   int64  `json:"roleId" v:"required" dc:"角色ID"`
	Avatar   string `json:"avatar" dc:"头像"`
	Sex      int    `json:"sex" v:"required" dc:"性别: 0未知 1男 2女"`
	Email    string `json:"email" dc:"邮箱"`
	Mobile   string `json:"mobile" dc:"手机号码"`
	Address  string `json:"address" dc:"联系地址"`
	Remark   string `json:"remark" dc:"备注"`
	Sort     int    `json:"sort" v:"required" dc:"排序"`
	Status   int    `json:"status" v:"required" dc:"状态: 0禁用 1启用"`
}

type CreateRes struct {
}

// UpdateReq 更新用户
type UpdateReq struct {
	g.Meta   `path:"/member/update" method:"put" tags:"用户管理" summary:"更新用户"`
	Uid      string `json:"uid" v:"required" dc:"用户ID"`
	Username string `json:"username" v:"required" dc:"用户名"`
	Nickname string `json:"nickname" v:"required" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
	Sex      int    `json:"sex" v:"required" dc:"性别: 0未知 1男 2女"`
	Email    string `json:"email" dc:"邮箱"`
	Mobile   string `json:"mobile" dc:"手机号码"`
	Address  string `json:"address" dc:"联系地址"`
	Remark   string `json:"remark" dc:"备注"`
	Sort     int    `json:"sort" v:"required" dc:"排序"`
	Status   int    `json:"status" v:"required" dc:"状态: 0禁用 1启用"`
	Password string `json:"password" dc:"密码（可选，不传则不更新密码）"`
	RoleId   int64  `json:"roleId" v:"required" dc:"角色ID"`
}

type UpdateRes struct {
}

// DeleteReq 删除用户
type DeleteReq struct {
	g.Meta `path:"/member/delete" method:"delete" tags:"用户管理" summary:"删除用户"`
	Uid    string `json:"uid" v:"required" dc:"用户ID"`
}

type DeleteRes struct {
}
