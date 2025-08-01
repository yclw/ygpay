// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"yclw/ygpay/api/user/v1"
	"yclw/ygpay/api/user/v2"
)

type IUserV1 interface {
	GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error)
	GetUserMenu(ctx context.Context, req *v1.GetUserMenuReq) (res *v1.GetUserMenuRes, err error)
}

type IUserV2 interface {
	GetUserInfo(ctx context.Context, req *v2.GetUserInfoReq) (res *v2.GetUserInfoRes, err error)
	GetUserMenu(ctx context.Context, req *v2.GetUserMenuReq) (res *v2.GetUserMenuRes, err error)
}
