package user

import (
	"cmp"
	"context"
	"slices"

	v1 "yclw/ygpay/api/user/v1"
	"yclw/ygpay/internal/logic/menu"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) GetUserMenu(ctx context.Context, req *v1.GetUserMenuReq) (res *v1.GetUserMenuRes, err error) {
	roleId := contexts.GetRoleId(ctx)

	// 获取角色菜单
	menus, err := c.MenuService.GetRoleMenu(ctx, roleId)
	if err != nil {
		return
	}

	// 转换为v2.UserMenu，并映射 id->UserMenu
	userMenus := make([]*v1.UserMenu, 0, len(menus))
	menuMap := make(map[int64]*v1.UserMenu)
	for _, menu := range menus {
		userMenu := c.menuModelToUserMenu(menu)
		userMenus = append(userMenus, userMenu)
		menuMap[menu.MenuInfo.Id] = userMenu
	}

	// 排序
	slices.SortFunc(userMenus, func(a, b *v1.UserMenu) int {
		return cmp.Compare(a.Meta.Rank, b.Meta.Rank)
	})

	// 构建菜单树
	tree := c.buildUserMenuTree(userMenus, menuMap)

	res = &v1.GetUserMenuRes{
		Menu: tree,
	}
	return
}

func (c *ControllerV1) menuModelToUserMenu(menuModel *menu.RoleMenuModel) *v1.UserMenu {
	userMenu := v1.UserMenu{
		Id:       menuModel.MenuInfo.Id,
		ParentId: menuModel.MenuInfo.Pid,
		Type:     menuModel.MenuInfo.Type,
		Name:     menuModel.MenuInfo.Name,
		Path:     menuModel.MenuInfo.Path,
		Meta: v1.UserMenuMeta{
			Icon:       menuModel.MenuInfo.Icon,
			Title:      menuModel.MenuInfo.Title,
			Rank:       int64(menuModel.MenuInfo.Sort),
			ShowParent: menuModel.MenuInfo.ShowParent == 1,
			KeepAlive:  menuModel.MenuInfo.KeepAlive == 1,
			ShowLink:   menuModel.MenuInfo.ShowLink == 1,
		},
	}
	switch menuModel.MenuInfo.Type {
	case menu.MenuTypeDir:
		userMenu.Redirect = menuModel.MenuInfo.Redirect
	case menu.MenuTypeMenu:
		userMenu.Component = menuModel.MenuInfo.Component
	case menu.MenuTypeLink:
		userMenu.Meta.FrameSrc = menuModel.MenuInfo.FrameSrc
		if menuModel.MenuInfo.Url != "" {
			userMenu.Name = menuModel.MenuInfo.Url
		}
	}
	return &userMenu
}

// 构建用户菜单树（userMenus和menuMap的内容其实是相同的，这里menuMap也可以用userMenus构建，但是需要遍历两次，第一次遍历构建map，第二次遍历构建树）
func (c *ControllerV1) buildUserMenuTree(userMenus []*v1.UserMenu, menuMap map[int64]*v1.UserMenu) (tree []*v1.UserMenu) {
	for _, node := range userMenus {
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
