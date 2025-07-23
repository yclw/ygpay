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
	g.Meta `path:"/member/create" method:"post" tags:"用户" summary:"创建用户"`
}

type CreateRes struct {
}

// UpdateReq 更新用户
type UpdateReq struct {
	g.Meta `path:"/member/update" method:"put" tags:"用户" summary:"更新用户"`
}

type UpdateRes struct {
}

// DeleteReq 删除用户
type DeleteReq struct {
	g.Meta `path:"/member/delete" method:"delete" tags:"用户" summary:"删除用户"`
}

type DeleteRes struct {
}
