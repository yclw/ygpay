// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package role

import (
	"context"

	"yclw/ygpay/api/role/v1"
)

type IRoleV1 interface {
	GetRoleApi(ctx context.Context, req *v1.GetRoleApiReq) (res *v1.GetRoleApiRes, err error)
	UpdateRoleApi(ctx context.Context, req *v1.UpdateRoleApiReq) (res *v1.UpdateRoleApiRes, err error)
	GetRoleMenu(ctx context.Context, req *v1.GetRoleMenuReq) (res *v1.GetRoleMenuRes, err error)
	UpdateRoleMenu(ctx context.Context, req *v1.UpdateRoleMenuReq) (res *v1.UpdateRoleMenuRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}
