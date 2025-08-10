// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package casbin

import (
	"context"

	"yclw/ygpay/api/casbin/v1"
)

type ICasbinV1 interface {
	SyncRoleApi(ctx context.Context, req *v1.SyncRoleApiReq) (res *v1.SyncRoleApiRes, err error)
	RefreshEnforcer(ctx context.Context, req *v1.RefreshEnforcerReq) (res *v1.RefreshEnforcerRes, err error)
}
