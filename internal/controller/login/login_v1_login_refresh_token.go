package login

import (
	"context"

	v1 "yclw/ygpay/api/login/v1"
)

func (c *ControllerV1) LoginRefreshToken(ctx context.Context, req *v1.LoginRefreshTokenReq) (res *v1.LoginRefreshTokenRes, err error) {
	res, err = c.LoginService.RefreshToken(ctx, req.RefreshToken)
	return
}
