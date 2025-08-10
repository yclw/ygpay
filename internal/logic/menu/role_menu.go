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

// UpdateRoleMenu 更新角色菜单
func (m *Menu) UpdateRoleMenu(ctx context.Context, roleId int64, menuUids []string) (err error) {
	// 获取菜单ID
	menuIds, err := dao.MenuInfo.FindIdsByMenuUids(ctx, menuUids)
	if err != nil {
		return
	}

	//TODO: 考虑事务

	// 删除角色菜单
	_, err = dao.RoleMenu.DeleteByRoleId(ctx, roleId)
	if err != nil {
		return
	}

	// 添加角色菜单
	_, err = dao.RoleMenu.AddRoleMenus(ctx, roleId, menuIds)
	return
}

// AddRoleMenu 添加角色菜单
func (m *Menu) AddRoleMenu(ctx context.Context, roleId int64, menuId int64) (err error) {
	_, err = dao.RoleMenu.AddRoleMenus(ctx, roleId, []int64{menuId})
	return
}
