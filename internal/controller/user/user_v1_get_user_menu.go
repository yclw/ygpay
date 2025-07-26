package user

import (
	"context"

	v1 "yclw/ygpay/api/user/v1"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) GetUserMenu(ctx context.Context, req *v1.GetUserMenuReq) (res *v1.GetUserMenuRes, err error) {
	roleId := contexts.GetRoleId(ctx)
	menus, err := c.MenuService.GetRoleMenu(ctx, roleId)
	if err != nil {
		return nil, err
	}
	res = &v1.GetUserMenuRes{
		Menu: menus,
	}
	return
}
