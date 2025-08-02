package menu

import (
	"context"

	v2 "yclw/ygpay/api/menu/v2"
	"yclw/ygpay/internal/model/entity"
)

func (c *ControllerV2) GetTree(ctx context.Context, req *v2.GetTreeReq) (res *v2.GetTreeRes, err error) {
	_, idTree, err := c.menuService.GetAllList(ctx)
	if err != nil {
		return
	}
	// 给idTree的每个节点添加title
	for _, node := range idTree.NodeMap {
		node.Title = node.Data.(*entity.MenuInfo).Title
	}

	// 将idTree转换为v2.MenuTree
	menuTree := make([]*v2.MenuTree, 0, len(idTree.Root.Children))
	menuTree = append(menuTree, idTree.Root.Children...)

	res = &v2.GetTreeRes{
		MenuTree: menuTree,
	}

	return
}
