package user

import (
	"context"

	v1 "yclw/ygpay/api/user/v1"
)

func (c *ControllerV1) GetUserMenu(ctx context.Context, req *v1.GetUserMenuReq) (res *v1.GetUserMenuRes, err error) {
	// roleId := contexts.GetRoleId(ctx)
	// _, idTree, err := c.MenuService.GetRoleMenu(ctx, roleId)
	if err != nil {
		return nil, err
	}
	res = &v1.GetUserMenuRes{
		// Menu: convertTreeToRoleMenu(idTree),
	}
	return
}

// // convertTreeToUserMenu 将IdTree转换为UserMenu切片
// func convertTreeToRoleMenu(idTree *tree.IdTree) []v1.UserMenu {
// 	if idTree == nil || idTree.Root == nil {
// 		return []v1.UserMenu{}
// 	}

// 	// 如果根节点ID是0，返回其子节点作为顶级菜单
// 	if idTree.Root.Id == 0 {
// 		result := make([]v1.UserMenu, 0, len(idTree.Root.Children))
// 		for _, child := range idTree.Root.Children {
// 			userMenu := convertTreeNodeToRoleMenu(child, idTree)
// 			result = append(result, userMenu)
// 		}
// 		return result
// 	}

// 	// 否则返回根节点本身
// 	userMenu := convertTreeNodeToRoleMenu(idTree.Root, idTree)
// 	return []v1.UserMenu{userMenu}
// }

// // convertTreeNodeToUserMenu 将TreeNode转换为UserMenu
// func convertTreeNodeToRoleMenu(node *tree.TreeNode, tree *tree.IdTree) v1.UserMenu {
// 	// 从Data中获取MenuInfo
// 	menuInfo, ok := node.Data.(*entity.MenuInfo)
// 	if !ok {
// 		// 如果转换失败，返回空的UserMenu
// 		return v1.UserMenu{}
// 	}

// 	// 递归转换子节点
// 	children := make([]v1.UserMenu, 0, len(node.Children))
// 	for _, child := range node.Children {
// 		childMenu := convertTreeNodeToRoleMenu(child, tree)
// 		children = append(children, childMenu)
// 	}

// 	// 获取路径列表（从根到当前节点的路径）
// 	pathNode := tree.GetPath(node.Id)
// 	pathList := make([]int64, 0, len(pathNode))
// 	for _, node := range pathNode {
// 		pathList = append(pathList, node.Id)
// 	}

// 	userMenu := v1.UserMenu{
// 		Name: menuInfo.Name,
// 		Path: menuInfo.Path,
// 		// Component:         menuInfo.Component,
// 		NoShowingChildren: gconv.Bool(menuInfo.NoShowingChildren),
// 		// Children:          children,
// 		Value: menuInfo.Value,
// 		Meta: v1.UserMenuMeta{
// 			Icon:       menuInfo.Icon,
// 			Title:      menuInfo.Title,
// 			Rank:       int64(menuInfo.Sort),
// 			ShowParent: gconv.Bool(menuInfo.ShowParent),
// 			// Auths:      []string{}, // 暂时为空数组，如果有权限控制需求可以在这里添加
// 		},
// 		ShowTooltip: gconv.Bool(menuInfo.ShowTooltip),
// 		// ParentId:    menuInfo.ParentId,
// 		// PathList: pathList,
// 		Redirect: menuInfo.Redirect,
// 	}
// 	if len(children) > 0 {
// 		userMenu.Children = children
// 	}
// 	if menuInfo.Redirect != "" {
// 		userMenu.Redirect = menuInfo.Redirect
// 	}
// 	if menuInfo.Component != "" {
// 		userMenu.Component = menuInfo.Component
// 	}
// 	return userMenu
// }
