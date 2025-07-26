package menu

import (
	"context"
	"fmt"
	v1 "yclw/ygpay/api/user/v1"
	"yclw/ygpay/internal/consts"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/model/entity"
	"yclw/ygpay/util/tree"

	"github.com/gogf/gf/v2/util/gconv"
)

var MenuService = NewMenu()

type Menu struct {
}

func NewMenu() *Menu {
	return &Menu{}
}

func (m *Menu) GetUserMenu(ctx context.Context, uid string) (res []v1.UserMenu, err error) {
	// 获取用户
	user, err := dao.MemberInfo.FindByUid(ctx, uid)
	if err != nil {
		return
	}
	// 获取用户角色
	role, err := dao.MemberRole.FindRoleIdByMemberId(ctx, user.Id)
	if err != nil {
		return
	}
	// 获取用户菜单
	menuIds, err := dao.RoleMenu.FindMenuIdsByRoleId(ctx, role)
	if err != nil {
		return
	}
	// 获取菜单
	menus, err := dao.MenuInfo.FindByMenuIds(ctx, menuIds)
	if err != nil {
		return
	}
	fmt.Println(menuIds)

	menuMap := make(map[int64]entity.MenuInfo)
	menuIds = make([]int64, 0)
	for _, menu := range menus {
		if menu.Status == consts.StatusDisabled {
			continue
		}
		menuMap[menu.Id] = menu
		menuIds = append(menuIds, menu.Id)
	}
	// 获得菜单ID树
	idTree, err := dao.MenuTree.BuildTreeByMenuIds(ctx, menuIds)
	if err != nil {
		return
	}
	for _, menu := range idTree.NodeMap {
		if menu.Id == 0 {
			continue
		}
		menu.Data = menuMap[menu.Id]
	}
	// 转换为UserMenu
	res = m.convertTreeToUserMenu(idTree)
	return
}

// convertTreeToUserMenu 将IdTree转换为UserMenu切片
func (m *Menu) convertTreeToUserMenu(idTree *tree.IdTree) []v1.UserMenu {
	if idTree == nil || idTree.Root == nil {
		return []v1.UserMenu{}
	}

	// 如果根节点ID是0，返回其子节点作为顶级菜单
	if idTree.Root.Id == 0 {
		result := make([]v1.UserMenu, 0, len(idTree.Root.Children))
		for _, child := range idTree.Root.Children {
			userMenu := m.convertTreeNodeToUserMenu(child, idTree)
			result = append(result, userMenu)
		}
		return result
	}

	// 否则返回根节点本身
	userMenu := m.convertTreeNodeToUserMenu(idTree.Root, idTree)
	return []v1.UserMenu{userMenu}
}

// convertTreeNodeToUserMenu 将TreeNode转换为UserMenu
func (m *Menu) convertTreeNodeToUserMenu(node *tree.TreeNode, tree *tree.IdTree) v1.UserMenu {
	// 从Data中获取MenuInfo
	menuInfo, ok := node.Data.(entity.MenuInfo)
	if !ok {
		// 如果转换失败，返回空的UserMenu
		return v1.UserMenu{}
	}

	// 递归转换子节点
	children := make([]v1.UserMenu, 0, len(node.Children))
	for _, child := range node.Children {
		childMenu := m.convertTreeNodeToUserMenu(child, tree)
		children = append(children, childMenu)
	}

	// 获取路径列表（从根到当前节点的路径）
	pathNode := tree.GetPath(node.Id)
	pathList := make([]int64, 0, len(pathNode))
	for _, node := range pathNode {
		pathList = append(pathList, node.Id)
	}

	userMenu := v1.UserMenu{
		Name: menuInfo.Name,
		Path: menuInfo.Path,
		// Component:         menuInfo.Component,
		NoShowingChildren: gconv.Bool(menuInfo.NoShowingChildren),
		// Children:          children,
		Value: menuInfo.Value,
		Meta: v1.UserMenuMeta{
			Icon:       menuInfo.Icon,
			Title:      menuInfo.Title,
			Rank:       int64(menuInfo.Sort),
			ShowParent: gconv.Bool(menuInfo.ShowParent),
			// Auths:      []string{}, // 暂时为空数组，如果有权限控制需求可以在这里添加
		},
		ShowTooltip: gconv.Bool(menuInfo.ShowTooltip),
		// ParentId:    menuInfo.ParentId,
		// PathList: pathList,
		Redirect: menuInfo.Redirect,
		// Redirect: "https://www.bilibili.com/video/BV1z7b9zrEqm/?t=22&spm_id_from=333.1007.tianma.5-2-16.click&vd_source=73cde154713a9a75e36905bb9b40ab16",
	}
	if len(children) > 0 {
		userMenu.Children = children
	}
	if menuInfo.Redirect != "" {
		userMenu.Redirect = menuInfo.Redirect
	}
	if menuInfo.Component != "" {
		userMenu.Component = menuInfo.Component
	}
	return userMenu
}
