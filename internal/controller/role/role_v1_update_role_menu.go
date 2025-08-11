package role

import (
	"context"
	"slices"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) UpdateRoleMenu(ctx context.Context, req *v1.UpdateRoleMenuReq) (res *v1.UpdateRoleMenuRes, err error) {
	operatorRoleUid := contexts.GetRoleUid(ctx)

	// 根据操作者RoleUid获取RoleId
	operatorRoleId, err := c.RoleService.GetRoleIdByUid(ctx, operatorRoleUid)
	if err != nil {
		return
	}

	// 获取可用菜单，当前为操作角色菜单
	enabledMenus, err := c.MenuService.GetRoleMenu(ctx, operatorRoleId)
	if err != nil {
		return
	}

	// 构建可用菜单map
	enabledMenusMap := make(map[string]bool)
	for _, menu := range enabledMenus {
		enabledMenusMap[menu.MenuInfo.MenuUid] = true
	}

	// 过滤掉无权限的菜单
	req.MenuList = slices.DeleteFunc(req.MenuList, func(menuUid string) bool {
		return !enabledMenusMap[menuUid]
	})

	// 根据目标roleUid获取roleId
	targetRoleId, err := c.RoleService.GetRoleIdByUid(ctx, req.RoleUid)
	if err != nil {
		return
	}

	err = c.MenuService.UpdateRoleMenu(ctx, targetRoleId, req.MenuList)
	return
}
