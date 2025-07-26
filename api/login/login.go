// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package login

import (
	"context"

	"yclw/ygpay/api/login/v1"
)

type ILoginV1 interface {
	LoginConfig(ctx context.Context, req *v1.LoginConfigReq) (res *v1.LoginConfigRes, err error)
	LoginCaptcha(ctx context.Context, req *v1.LoginCaptchaReq) (res *v1.LoginCaptchaRes, err error)
	AccountLogin(ctx context.Context, req *v1.AccountLoginReq) (res *v1.AccountLoginRes, err error)
	LoginLogout(ctx context.Context, req *v1.LoginLogoutReq) (res *v1.LoginLogoutRes, err error)
	LoginRefreshToken(ctx context.Context, req *v1.LoginRefreshTokenReq) (res *v1.LoginRefreshTokenRes, err error)
}
