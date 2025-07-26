package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ApiModel struct {
	Id          int64
	Name        string
	Path        string
	Method      string
	GroupName   string
	Description string
	NeedAuth    int
	RateLimit   int
	Sort        int
	Status      int
	CreatedAt   *gtime.Time
	UpdatedAt   *gtime.Time
}

// GetListReq 获取用户列表
type GetListReq struct {
	g.Meta `path:"/api/list" method:"get" tags:"API管理" summary:"获取API列表"`
}

type GetListRes struct {
	List []*ApiModel
}

// GetOneReq 获取用户详情
type GetOneReq struct {
	g.Meta `path:"/api/one" method:"get" tags:"API管理" summary:"获取API详情"`
}

type GetOneRes struct {
	Member *ApiModel
}

// CreateReq 创建用户
type CreateReq struct {
	g.Meta      `path:"/api/create" method:"post" tags:"API管理" summary:"创建API"`
	Name        string `json:"name" v:"required" dc:"API名称"`
	Path        string `json:"path" v:"required" dc:"API路径"`
	Method      string `json:"method" v:"required" dc:"API方法"`
	GroupName   string `json:"groupName" v:"required" dc:"API分组"`
	Description string `json:"description" v:"required" dc:"API描述"`
	NeedAuth    int    `json:"needAuth" v:"required" dc:"是否需要认证"`
	RateLimit   int    `json:"rateLimit" v:"required" dc:"限流次数/分钟"`
	Sort        int    `json:"sort" v:"required" dc:"排序"`
	Status      int    `json:"status" v:"required" dc:"状态"`
}

type CreateRes struct {
}

// UpdateReq 更新用户
type UpdateReq struct {
	g.Meta      `path:"/api/update" method:"put" tags:"API管理" summary:"更新API"`
	Id          int64  `json:"id" v:"required" dc:"APIID"`
	Name        string `json:"name" v:"required" dc:"API名称"`
	Path        string `json:"path" v:"required" dc:"API路径"`
	Method      string `json:"method" v:"required" dc:"API方法"`
	GroupName   string `json:"groupName" v:"required" dc:"API分组"`
	Description string `json:"description" v:"required" dc:"API描述"`
	NeedAuth    int    `json:"needAuth" v:"required" dc:"是否需要认证"`
	RateLimit   int    `json:"rateLimit" v:"required" dc:"限流次数/分钟"`
	Sort        int    `json:"sort" v:"required" dc:"排序"`
	Status      int    `json:"status" v:"required" dc:"状态"`
}

type UpdateRes struct {
}

// DeleteReq 删除用户
type DeleteReq struct {
	g.Meta `path:"/api/delete" method:"delete" tags:"API管理" summary:"删除API"`
}

type DeleteRes struct {
}
