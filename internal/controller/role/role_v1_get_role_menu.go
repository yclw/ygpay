package role

import (
	"context"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/internal/logic/menu"
	"yclw/ygpay/util/tree"
)

func (c *ControllerV1) GetRoleMenu(ctx context.Context, req *v1.GetRoleMenuReq) (res *v1.GetRoleMenuRes, err error) {

	// 获取角色可用菜单
	_, idTree, err := c.MenuService.GetAllList(ctx)
	if err != nil {
		return
	}

	// 获取角色使用菜单
	_, useIdTree, err := c.MenuService.GetRoleMenu(ctx, req.Id)
	if err != nil {
		return
	}
	res = &v1.GetRoleMenuRes{
		Tree: c.roleMenuModelToV1Trees(idTree, useIdTree),
	}

	return
}

func (c *ControllerV1) roleMenuModelToV1Trees(idTree *tree.IdTree, useIdTree *tree.IdTree) []*v1.RoleMenuTreeModel {

	useMenuMap := useIdTree.NodeMap

	res := make([]*v1.RoleMenuTreeModel, 0, len(idTree.Root.Children))
	for _, node := range idTree.Root.Children {
		res = append(res, c.roleMenuModelToV1Tree(node, useMenuMap))
	}
	return res
}

// RoleModelToV1Tree 递归构建RoleTreeModel
func (c *ControllerV1) roleMenuModelToV1Tree(node *tree.TreeNode, useMenuMap map[int64]*tree.TreeNode) *v1.RoleMenuTreeModel {
	// 构建RoleTreeModel
	roleModel := node.Data.(*menu.MenuModel)
	_, exist := useMenuMap[roleModel.MenuInfo.Id]
	roleTreeModel := &v1.RoleMenuTreeModel{
		RoleMenuModel: &v1.RoleMenuModel{
			Id:   roleModel.MenuInfo.Id,
			Name: roleModel.MenuInfo.Name,
			Sort: roleModel.MenuInfo.Sort,
			Use:  exist,
		},
	}
	// 递归构建子节点
	for _, child := range node.Children {
		roleTreeModel.Children = append(roleTreeModel.Children, c.roleMenuModelToV1Tree(child, useMenuMap))
	}
	return roleTreeModel
}
