package menu

import (
	"context"
	"yclw/ygpay/internal/dao"
)

// GetRoleMenu 获取角色菜单
func (m *Menu) GetRoleMenu(ctx context.Context, roleId int64) (res []*RoleMenuModel, err error) {
	// 获取角色菜单ID
	menuIds, err := dao.RoleMenu.FindMenuIdsByRoleId(ctx, roleId)
	if err != nil {
		return
	}

	// 获取角色菜单信息
	menus, err := dao.MenuInfo.FindEnabledByMenuIds(ctx, menuIds)
	if err != nil {
		return
	}

	// 转换为RoleMenuModel
	res = make([]*RoleMenuModel, 0, len(menuIds))
	for _, menu := range menus {
		res = append(res, &RoleMenuModel{MenuInfo: menu})
	}

	return
}
