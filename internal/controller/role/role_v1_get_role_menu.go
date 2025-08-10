package role

import (
	"cmp"
	"context"
	"slices"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/internal/logic/menu"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) GetRoleMenu(ctx context.Context, req *v1.GetRoleMenuReq) (res *v1.GetRoleMenuRes, err error) {
	operator := contexts.GetRoleId(ctx)

	// 获取可用菜单，当前为操作角色菜单
	enabledMenus, err := c.MenuService.GetRoleMenu(ctx, operator)
	if err != nil {
		return
	}

	// 获取已用菜单
	usedMenus, err := c.MenuService.GetRoleMenu(ctx, req.Id)
	if err != nil {
		return
	}

	// 构建已用菜单map
	usedMenusMap := make(map[int64]bool)
	for _, menu := range usedMenus {
		usedMenusMap[menu.MenuInfo.Id] = true
	}

	// 转换可用菜单为v1.MenuModel，并构建可用菜单map
	enabledMenusMap := make(map[int64]*v1.MenuModel)
	roleMenus := make([]*v1.MenuModel, 0)
	for _, menu := range enabledMenus {
		roleMenu := c.menuModelToRoleMenu(menu, usedMenusMap)
		roleMenus = append(roleMenus, roleMenu)
		enabledMenusMap[menu.MenuInfo.Id] = roleMenu
	}

	// 排序
	slices.SortFunc(roleMenus, func(a, b *v1.MenuModel) int {
		return cmp.Compare(a.Sort, b.Sort)
	})

	// 构建菜单树
	tree := c.buildRoleMenuTree(roleMenus, enabledMenusMap)

	res = &v1.GetRoleMenuRes{
		Tree: tree,
	}

	return
}

func (c *ControllerV1) menuModelToRoleMenu(menuModel *menu.RoleMenuModel, usedMenusMap map[int64]bool) *v1.MenuModel {
	roleMenu := v1.MenuModel{
		MenuUid:  menuModel.MenuInfo.MenuUid,
		ParentId: menuModel.MenuInfo.Pid,
		Title:    menuModel.MenuInfo.Title,
		Sort:     menuModel.MenuInfo.Sort,
		Use:      usedMenusMap[menuModel.MenuInfo.Id],
	}
	return &roleMenu
}

func (c *ControllerV1) buildRoleMenuTree(roleMenus []*v1.MenuModel, menuMap map[int64]*v1.MenuModel) (tree []*v1.MenuModel) {
	for _, node := range roleMenus {
		parentId := node.ParentId
		// 查找父节点
		if parent, exists := menuMap[parentId]; exists {
			// 将当前节点添加到父节点的Children中
			parent.Children = append(parent.Children, node)
		} else if parentId == 0 {
			// 无父节点，作为根节点
			tree = append(tree, node)
		}
	}
	return
}
