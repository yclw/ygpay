// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"yclw/ygpay/api/user/v1"
)

type IUserV1 interface {
	UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error)
	LoginConfig(ctx context.Context, req *v1.LoginConfigReq) (res *v1.LoginConfigRes, err error)
	LoginCaptcha(ctx context.Context, req *v1.LoginCaptchaReq) (res *v1.LoginCaptchaRes, err error)
	AccountLogin(ctx context.Context, req *v1.AccountLoginReq) (res *v1.AccountLoginRes, err error)
	LoginLogout(ctx context.Context, req *v1.LoginLogoutReq) (res *v1.LoginLogoutRes, err error)
}
