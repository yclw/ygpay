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
