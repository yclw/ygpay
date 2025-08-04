package role

import (
	"context"
	"slices"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) UpdateRoleMenu(ctx context.Context, req *v1.UpdateRoleMenuReq) (res *v1.UpdateRoleMenuRes, err error) {
	operator := contexts.GetRoleId(ctx)

	// 获取可用菜单，当前为操作角色菜单
	enabledMenus, err := c.MenuService.GetRoleMenu(ctx, operator)
	if err != nil {
		return
	}

	// 构建可用菜单map
	enabledMenusMap := make(map[int64]bool)
	for _, menu := range enabledMenus {
		enabledMenusMap[menu.MenuInfo.Id] = true
	}

	// 过滤掉无权限的菜单
	req.MenuList = slices.DeleteFunc(req.MenuList, func(menuId int64) bool {
		return !enabledMenusMap[menuId]
	})

	err = c.MenuService.UpdateRoleMenu(ctx, req.Id, req.MenuList)
	return
}
