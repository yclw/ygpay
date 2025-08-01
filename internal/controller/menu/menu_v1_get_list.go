package menu

import (
	"context"

	v1 "yclw/ygpay/api/menu/v1"
	"yclw/ygpay/util/tree"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	_, idTree, err := c.menuService.GetAllList(ctx)
	if err != nil {
		return
	}
	res = &v1.GetListRes{}
	res.Tree = c.menuModelToV1Trees(idTree)
	return
}

// RoleModelToV1Trees 构建RoleTreeModel列表
func (c *ControllerV1) menuModelToV1Trees(idTree *tree.IdTree) []*v1.MenuTreeModel {
	res := make([]*v1.MenuTreeModel, 0, len(idTree.Root.Children))
	for _, node := range idTree.Root.Children {
		res = append(res, c.menuModelToV1Tree(node))
	}
	return res
}

// RoleModelToV1Tree 递归构建RoleTreeModel
func (c *ControllerV1) menuModelToV1Tree(node *tree.TreeNode) *v1.MenuTreeModel {
	// 构建RoleTreeModel
	// menuModel := node.Data.(*menu.MenuModel)
	menuTreeModel := &v1.MenuTreeModel{
		MenuModel: &v1.MenuModel{
			// Id:                menuModel.MenuInfo.Id,
			// Pid:               node.Pid,
			// Name:              menuModel.MenuInfo.Name,
			// Path:              menuModel.MenuInfo.Path,
			// Icon:              menuModel.MenuInfo.Icon,
			// Title:             menuModel.MenuInfo.Title,
			// ShowParent:        menuModel.MenuInfo.ShowParent,
			// Component:         menuModel.MenuInfo.Component,
			// NoShowingChildren: menuModel.MenuInfo.NoShowingChildren,
			// Value:             menuModel.MenuInfo.Value,
			// ShowTooltip:       menuModel.MenuInfo.ShowTooltip,
			// ParentId:          menuModel.MenuInfo.ParentId,
			// Redirect:          menuModel.MenuInfo.Redirect,
			// Description:       menuModel.MenuInfo.Description,
			// Sort:              menuModel.MenuInfo.Sort,
			// Status:            menuModel.MenuInfo.Status,
			// CreatedAt:         menuModel.MenuInfo.CreatedAt,
			// UpdatedAt:         menuModel.MenuInfo.UpdatedAt,
		},
	}
	// 递归构建子节点
	for _, child := range node.Children {
		menuTreeModel.Children = append(menuTreeModel.Children, c.menuModelToV1Tree(child))
	}
	return menuTreeModel
}
