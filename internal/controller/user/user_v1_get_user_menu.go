package user

import (
	"context"

	v1 "yclw/ygpay/api/user/v1"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) GetUserMenu(ctx context.Context, req *v1.GetUserMenuReq) (res *v1.GetUserMenuRes, err error) {
	uid := contexts.GetUserUid(ctx)
	menus, err := c.MenuService.GetUserMenu(ctx, uid)
	if err != nil {
		return nil, err
	}
	res = &v1.GetUserMenuRes{
		Menu: menus,
	}
	return
}
