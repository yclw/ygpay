// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package menu

import (
	"context"

	"yclw/ygpay/api/menu/v1"
	"yclw/ygpay/api/menu/v2"
)

type IMenuV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}

type IMenuV2 interface {
	GetList(ctx context.Context, req *v2.GetListReq) (res *v2.GetListRes, err error)
	GetOne(ctx context.Context, req *v2.GetOneReq) (res *v2.GetOneRes, err error)
	Create(ctx context.Context, req *v2.CreateReq) (res *v2.CreateRes, err error)
	Update(ctx context.Context, req *v2.UpdateReq) (res *v2.UpdateRes, err error)
	Delete(ctx context.Context, req *v2.DeleteReq) (res *v2.DeleteRes, err error)
	GetTree(ctx context.Context, req *v2.GetTreeReq) (res *v2.GetTreeRes, err error)
}
