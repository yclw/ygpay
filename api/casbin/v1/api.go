package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SyncRoleApiReq 同步角色API到Casbin
type SyncRoleApiReq struct {
	g.Meta `path:"/casbin/sync" method:"post" tags:"Casbin管理" summary:"同步角色API到Casbin"`
}

type SyncRoleApiRes struct {
}

// RefreshEnforcerReq 刷新Casbin Enforcer
type RefreshEnforcerReq struct {
	g.Meta `path:"/casbin/refresh" method:"post" tags:"Casbin管理" summary:"刷新Casbin Enforcer"`
}

type RefreshEnforcerRes struct {
}

// GetPoliciesReq 获取Casbin策略列表
type GetPoliciesReq struct {
	g.Meta `path:"/casbin/policies" method:"get" tags:"Casbin管理" summary:"获取Casbin策略列表"`
	RoleId *int64 `json:"roleId" dc:"角色ID（可选，不传则获取所有策略）"`
}

type GetPoliciesRes struct {
	Policies []PolicyModel `json:"policies" dc:"策略列表"`
}

// PolicyModel Casbin策略模型
type PolicyModel struct {
	RoleId string `json:"roleId" dc:"角色ID"`
	Path   string `json:"path" dc:"API路径"`
	Method string `json:"method" dc:"请求方法"`
}

// AddPolicyReq 添加Casbin策略
type AddPolicyReq struct {
	g.Meta `path:"/casbin/policy/add" method:"post" tags:"Casbin管理" summary:"添加Casbin策略"`
	RoleId string `json:"roleId" v:"required" dc:"角色ID"`
	Path   string `json:"path" v:"required" dc:"API路径"`
	Method string `json:"method" v:"required" dc:"请求方法"`
}

type AddPolicyRes struct {
}

// RemovePolicyReq 删除Casbin策略
type RemovePolicyReq struct {
	g.Meta `path:"/casbin/policy/remove" method:"delete" tags:"Casbin管理" summary:"删除Casbin策略"`
	RoleId string `json:"roleId" v:"required" dc:"角色ID"`
	Path   string `json:"path" v:"required" dc:"API路径"`
	Method string `json:"method" v:"required" dc:"请求方法"`
}

type RemovePolicyRes struct {
}

// GetApisByRoleReq 获取角色对应的API列表（从Casbin中）
type GetApisByRoleReq struct {
	g.Meta `path:"/casbin/role/apis" method:"get" tags:"Casbin管理" summary:"获取角色对应的API列表"`
	RoleId int64 `json:"roleId" v:"required" dc:"角色ID"`
}

type GetApisByRoleRes struct {
	Apis []ApiModel `json:"apis" dc:"API列表"`
}

// ApiModel API模型
type ApiModel struct {
	Path   string `json:"path" dc:"API路径"`
	Method string `json:"method" dc:"请求方法"`
}
