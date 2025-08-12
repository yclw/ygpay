package menu

import (
	"context"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/util/slices"
)

// GetRoleMenu 获取角色菜单分配数据（带Use标记的菜单列表）
func (m *Menu) GetRoleMenu(ctx context.Context, operatorRoleUid, targetRoleUid string) (res []*RoleMenuModel, err error) {

	// 获取操作者角色ID可用的菜单列表
	operatorRoleId, err := dao.RoleInfo.FindIdByRoleUid(ctx, operatorRoleUid)
	if err != nil {
		return nil, err
	}
	enabledMenuIds, err := dao.RoleMenu.FindMenuIdsByRoleId(ctx, operatorRoleId)
	if err != nil {
		return nil, err
	}
	enabledMenus, err := dao.MenuInfo.FindEnabledByMenuIds(ctx, enabledMenuIds)
	if err != nil {
		return nil, err
	}

	// 获取目标角色已用的菜单ID列表
	usedMenuIds := enabledMenuIds
	if operatorRoleUid != targetRoleUid {
		// 获取目标角色已用的菜单ID列表
		targetRoleId, err := dao.RoleInfo.FindIdByRoleUid(ctx, targetRoleUid)
		if err != nil {
			return nil, err
		}
		usedMenuIds, err = dao.RoleMenu.FindMenuIdsByRoleId(ctx, targetRoleId)
		if err != nil {
			return nil, err
		}
	}

	usedMenuMap := make(map[int64]bool, len(usedMenuIds))
	for _, id := range usedMenuIds {
		usedMenuMap[id] = true
	}

	// 转换为RoleMenuModel
	res = make([]*RoleMenuModel, 0, len(enabledMenus))
	for _, menu := range enabledMenus {
		res = append(res, &RoleMenuModel{
			MenuInfo: menu,
			Use:      usedMenuMap[menu.Id],
		})
	}

	return
}

// UpdateRoleMenu 更新角色菜单（重构：在Logic层处理权限检查）
func (m *Menu) UpdateRoleMenu(ctx context.Context, operatorRoleUid, targetRoleUid string, menuUids []string) (err error) {

	// 获取目标角色ID
	targetRoleId, err := dao.RoleInfo.FindIdByRoleUid(ctx, targetRoleUid)
	if err != nil {
		return
	}

	// 将菜单UID列表转换为ID列表
	menuIds, err := dao.MenuInfo.FindIdsByMenuUids(ctx, menuUids)
	if err != nil {
		return
	}

	// 权限过滤：只允许操作者有权限的菜单
	operatorRoleId, err := dao.RoleInfo.FindIdByRoleUid(ctx, operatorRoleUid)
	if err != nil {
		return
	}
	filteredMenuIds, err := m.FilterAllowedMenuIdsByRole(ctx, menuIds, operatorRoleId)
	if err != nil {
		return
	}

	// 更新角色菜单
	_, err = dao.RoleMenu.UpdateRoleMenu(ctx, targetRoleId, filteredMenuIds)
	return
}

// FilterAllowedMenuIdsByRole 根据角色过滤允许的菜单列表，直接返回ID
func (m *Menu) FilterAllowedMenuIdsByRole(ctx context.Context, menuIds []int64, roleId int64) ([]int64, error) {

	// 获取角色允许的菜单ID列表
	allowedMenuIds, err := dao.RoleMenu.FindMenuIdsByRoleId(ctx, roleId)
	if err != nil {
		return nil, err
	}

	// 计算交集（只保留既在请求列表中又被角色允许的菜单）
	return slices.IntersectInt64s(menuIds, allowedMenuIds), nil
}
