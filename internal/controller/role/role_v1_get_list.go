package role

import (
	"cmp"
	"context"
	"slices"

	v1 "yclw/ygpay/api/role/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {

	// 获取所有角色
	roles, err := c.RoleService.GetAllList(ctx)
	if err != nil {
		return
	}

	// 转换为v1.RoleModel，并映射 id->RoleModel
	models := make([]*v1.RoleModel, 0, len(roles))
	roleMap := make(map[int64]*v1.RoleModel)
	for _, role := range roles {
		roleModel := c.roleModelToV1(role)
		models = append(models, roleModel)
		roleMap[role.Id] = roleModel
	}

	// 排序
	slices.SortFunc(models, func(a, b *v1.RoleModel) int {
		return cmp.Compare(a.Sort, b.Sort)
	})

	// 构建角色树
	tree := buildRoleTree(models, roleMap)

	// 构建响应
	res = &v1.GetListRes{
		List: models,
		Tree: tree,
	}

	return
}

// buildRoleTree 构建角色树
func buildRoleTree(roles []*v1.RoleModel, roleMap map[int64]*v1.RoleModel) (tree []*v1.RoleModel) {
	for _, node := range roles {
		parentId := node.ParentId
		// 查找父节点
		if parent, exists := roleMap[parentId]; exists {
			// 将当前节点添加到父节点的Children中
			parent.Children = append(parent.Children, node)
		} else {
			// 无父节点，作为根节点
			tree = append(tree, node)
		}
	}
	return tree
}
