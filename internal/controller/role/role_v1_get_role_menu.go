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
	// 获取操作者角色UID
	operatorRoleUid := contexts.GetRoleUid(ctx)

	// 获得带Use标记的菜单列表
	enabledMenus, err := c.MenuService.GetRoleMenu(ctx, operatorRoleUid, req.RoleUid)
	if err != nil {
		return
	}

	// 转换为响应格式并构建菜单映射
	roleMenus, menuMap := c.convertToMenuModels(enabledMenus)
	c.sortMenuModels(roleMenus)

	// 构建菜单树
	tree := c.buildRoleMenuTree(roleMenus, menuMap)

	res = &v1.GetRoleMenuRes{
		Tree: tree,
	}
	return
}

// 转换菜单模型为响应格式并返回映射表
func (c *ControllerV1) convertToMenuModels(menus []*menu.RoleMenuModel) ([]*v1.MenuModel, map[int64]*v1.MenuModel) {
	roleMenus := make([]*v1.MenuModel, 0, len(menus))
	menuMap := make(map[int64]*v1.MenuModel, len(menus))

	for _, menu := range menus {
		roleMenu := &v1.MenuModel{
			MenuUid:  menu.MenuInfo.MenuUid,
			ParentId: menu.MenuInfo.Pid,
			Title:    menu.MenuInfo.Title,
			Sort:     menu.MenuInfo.Sort,
			Use:      menu.Use,
		}
		roleMenus = append(roleMenus, roleMenu)
		menuMap[menu.MenuInfo.Id] = roleMenu
	}
	return roleMenus, menuMap
}

func (c *ControllerV1) sortMenuModels(models []*v1.MenuModel) {
	slices.SortFunc(models, func(a, b *v1.MenuModel) int {
		return cmp.Compare(a.Sort, b.Sort)
	})
}

// 菜单树构建
func (c *ControllerV1) buildRoleMenuTree(roleMenus []*v1.MenuModel, menuMap map[int64]*v1.MenuModel) []*v1.MenuModel {
	tree := make([]*v1.MenuModel, 0)
	for _, node := range roleMenus {
		parentId := node.ParentId
		if parent, exists := menuMap[parentId]; exists {
			parent.Children = append(parent.Children, node)
		} else if parentId == 0 {
			tree = append(tree, node)
		}
	}
	return tree
}
