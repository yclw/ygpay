// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package casbin

import (
	"context"

	v1 "yclw/ygpay/api/casbin/v1"
)

type ICasbinV1 interface {
	SyncRoleApi(ctx context.Context, req *v1.SyncRoleApiReq) (res *v1.SyncRoleApiRes, err error)
	RefreshEnforcer(ctx context.Context, req *v1.RefreshEnforcerReq) (res *v1.RefreshEnforcerRes, err error)
	GetPolicies(ctx context.Context, req *v1.GetPoliciesReq) (res *v1.GetPoliciesRes, err error)
	AddPolicy(ctx context.Context, req *v1.AddPolicyReq) (res *v1.AddPolicyRes, err error)
	RemovePolicy(ctx context.Context, req *v1.RemovePolicyReq) (res *v1.RemovePolicyRes, err error)
	GetApisByRole(ctx context.Context, req *v1.GetApisByRoleReq) (res *v1.GetApisByRoleRes, err error)
}
