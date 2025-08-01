package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ApiModel struct {
	Id          int64       `json:"id" dc:"APIID"`
	Name        string      `json:"name" dc:"API名称"`
	Path        string      `json:"path" dc:"API路径"`
	Method      string      `json:"method" dc:"API方法"`
	GroupName   string      `json:"groupName" dc:"API分组"`
	Description string      `json:"description" dc:"API描述"`
	NeedAuth    int         `json:"needAuth" dc:"是否需要认证"`
	RateLimit   int         `json:"rateLimit" dc:"限流次数/分钟"`
	Sort        int         `json:"sort" dc:"排序"`
	Status      int         `json:"status" dc:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// GetListReq 获取API列表
type GetListReq struct {
	g.Meta    `path:"/api/list" method:"get" tags:"API管理" summary:"获取API列表"`
	Page      int         `json:"page" v:"required|min:1" dc:"页码"`
	Size      int         `json:"size" v:"required|between:1,100" dc:"每页条数"`
	NeedAuth  *int        `json:"needAuth" dc:"是否需要认证（0:否 1:是）"`
	StartDate *gtime.Time `json:"startDate" dc:"开始日期"`
	EndDate   *gtime.Time `json:"endDate" dc:"结束日期"`
	SortField string      `json:"sortField" dc:"排序字段"`
	SortDesc  bool        `json:"sortDesc" dc:"是否降序"`
	Name      string      `json:"name" dc:"API名称"`
	Path      string      `json:"path" dc:"API路径"`
	Method    string      `json:"method" dc:"API方法"`
	GroupName string      `json:"groupName" dc:"API分组"`
	Status    *int        `json:"status" dc:"状态筛选（0:禁用 1:启用）"`
}

type GetListRes struct {
	List  []*ApiModel `json:"list" dc:"API列表"`
	Total int         `json:"total" dc:"总条数"`
}

// GetOneReq 获取API详情
type GetOneReq struct {
	g.Meta `path:"/api/one" method:"get" tags:"API管理" summary:"获取API详情"`
	Id     int64 `json:"id" v:"required" dc:"APIID"`
}

type GetOneRes struct {
	*ApiModel
}

// CreateReq 创建API
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

// UpdateReq 更新API
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

// DeleteReq 删除API
type DeleteReq struct {
	g.Meta `path:"/api/delete" method:"delete" tags:"API管理" summary:"删除API"`
	Id     int64 `json:"id" v:"required" dc:"APIID"`
}

type DeleteRes struct {
}
