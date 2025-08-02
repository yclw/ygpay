package user

import (
	"context"
	"sort"

	v2 "yclw/ygpay/api/user/v2"
	"yclw/ygpay/internal/logic/menu"
	"yclw/ygpay/internal/model/entity"
	"yclw/ygpay/pkg/contexts"

	"yclw/ygpay/util/tree"
)

func (c *ControllerV2) GetUserMenu(ctx context.Context, req *v2.GetUserMenuReq) (res *v2.GetUserMenuRes, err error) {
	roleId := contexts.GetRoleId(ctx)
	_, idTree, err := c.MenuService.GetRoleMenu(ctx, roleId)
	if err != nil {
		return
	}
	res = &v2.GetUserMenuRes{
		Menu: c.treeToRoleMenu(idTree),
	}
	return
}

// 将树形结构转换为菜单结构
func (c *ControllerV2) treeToRoleMenu(idTree *tree.IdTree) []v2.UserMenu {
	menus := make([]v2.UserMenu, 0)
	for _, node := range idTree.Root.Children {
		menu := c.nodeToUserMenu(node)
		menus = append(menus, menu)
	}
	// 排序
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].Meta.Rank > menus[j].Meta.Rank
	})
	return menus
}

// 递归转换树形结构节点为菜单结构
func (c *ControllerV2) nodeToUserMenu(node *tree.TreeNode) v2.UserMenu {
	menuModel := node.Data.(*entity.MenuInfo)
	userMenu := v2.UserMenu{
		Type: menuModel.Type,
		Name: menuModel.Name,
		Path: menuModel.Path,
		Meta: v2.UserMenuMeta{
			Icon:       menuModel.Icon,
			Title:      menuModel.Title,
			Rank:       int64(menuModel.Sort),
			ShowParent: menuModel.ShowParent == 1,
			KeepAlive:  menuModel.KeepAlive == 1,
			ShowLink:   menuModel.ShowLink == 1,
		},
	}
	switch menuModel.Type {
	case menu.MenuTypeDir:
		userMenu.Redirect = menuModel.Redirect
	case menu.MenuTypeMenu:
		userMenu.Component = menuModel.Component
	case menu.MenuTypeLink:
		userMenu.FrameSrc = menuModel.FrameSrc
	}
	for _, child := range node.Children {
		childMenu := c.nodeToUserMenu(child)
		userMenu.Children = append(userMenu.Children, childMenu)
	}

	// 排序
	sort.Slice(userMenu.Children, func(i, j int) bool {
		return userMenu.Children[i].Meta.Rank > userMenu.Children[j].Meta.Rank
	})

	return userMenu
}
