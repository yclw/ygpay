package menu

import (
	"context"

	v1 "yclw/ygpay/api/menu/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	// 获取所有菜单
	models, err := c.menuService.GetAllList(ctx)
	if err != nil {
		return
	}

	// 转换为v2.MenuModel，并映射 id->MenuModel
	menus := make([]*v1.MenuModel, 0, len(models))
	menuMap := make(map[int64]*v1.MenuModel)
	for _, model := range models {
		menu := c.menuModelToV1(model)
		menus = append(menus, menu)
		menuMap[model.Id] = menu
	}

	// 构建菜单树
	tree := buildMenuTree(menus, menuMap)

	// 构建响应
	res = &v1.GetListRes{
		List: tree,
		Tree: menus,
	}
	return
}

// buildMenuTree 构建菜单树
func buildMenuTree(menus []*v1.MenuModel, menuMap map[int64]*v1.MenuModel) (tree []*v1.MenuModel) {
	for _, node := range menus {
		parentId := node.ParentId
		// 查找父节点
		if parent, exists := menuMap[parentId]; exists {
			// 将当前节点添加到父节点的Children中
			parent.Children = append(parent.Children, node)
		} else {
			// 无父节点，作为根节点
			tree = append(tree, node)
		}
	}
	return
}
