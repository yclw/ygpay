package login

import (
	"context"

	v1 "yclw/ygpay/api/login/v1"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) LoginLogout(ctx context.Context, req *v1.LoginLogoutReq) (res *v1.LoginLogoutRes, err error) {
	token := g.RequestFromCtx(ctx).GetHeader("Authorization")
	err = c.LoginService.Logout(ctx, token)
	return
}
