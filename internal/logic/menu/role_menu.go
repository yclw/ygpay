package menu

import (
	"context"
	"yclw/ygpay/internal/consts"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/model/entity"
	"yclw/ygpay/util/tree"
)

// GetRoleMenu 获取角色菜单
func (m *Menu) GetRoleMenu(ctx context.Context, roleId int64) (res []*RoleMenuModel, idTree *tree.IdTree, err error) {
	// 获取角色菜单ID
	menuIds, err := dao.RoleMenu.FindMenuIdsByRoleId(ctx, roleId)
	if err != nil {
		return
	}

	// 获取角色菜单信息
	menus, err := dao.MenuInfo.FindByMenuIds(ctx, menuIds)
	if err != nil {
		return
	}

	// 去除禁用菜单
	menuMap := make(map[int64]*entity.MenuInfo)
	menuIds = make([]int64, 0)
	for _, menu := range menus {
		if menu.Status == consts.StatusDisabled {
			continue
		}
		menuMap[menu.Id] = menu
		menuIds = append(menuIds, menu.Id)
	}

	// 获取菜单ID树关系
	Ts, err := dao.MenuTree.FindByMenuIds(ctx, menuIds)
	if err != nil {
		return
	}

	// 构建菜单树
	Ts = append(Ts, tree.T{Id: 0, Pid: -1})
	idTree, err = tree.BuildTree(Ts, 0)
	if err != nil {
		return
	}
	for _, menu := range idTree.NodeMap {
		if menu.Id == 0 {
			continue
		}
		menu.Data = menuMap[menu.Id]
	}

	// 转换为RoleMenuModel
	res = make([]*RoleMenuModel, 0, len(menus))
	for _, menuId := range menuIds {
		res = append(res, &RoleMenuModel{MenuInfo: menuMap[menuId], TreeNode: idTree.NodeMap[menuId]})
	}

	return
}
