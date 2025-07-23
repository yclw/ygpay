package user

import (
	"context"

	v1 "yclw/ygpay/api/user/v1"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) LoginLogout(ctx context.Context, _ *v1.LoginLogoutReq) (_ *v1.LoginLogoutRes, err error) {
	token := g.RequestFromCtx(ctx).GetHeader("Authorization")
	err = c.LoginService.Logout(ctx, token)
	return
}
