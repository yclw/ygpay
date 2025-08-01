// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package role

import (
	"context"

	"yclw/ygpay/api/role/v1"
	"yclw/ygpay/api/role/v2"
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

type IRoleV2 interface {
	GetList(ctx context.Context, req *v2.GetListReq) (res *v2.GetListRes, err error)
	GetOne(ctx context.Context, req *v2.GetOneReq) (res *v2.GetOneRes, err error)
	Create(ctx context.Context, req *v2.CreateReq) (res *v2.CreateRes, err error)
	Update(ctx context.Context, req *v2.UpdateReq) (res *v2.UpdateRes, err error)
	Delete(ctx context.Context, req *v2.DeleteReq) (res *v2.DeleteRes, err error)
}
