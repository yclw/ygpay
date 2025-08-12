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

	// 转换为v1.RoleModel，并映射 roleId->RoleModel
	models := make([]*v1.RoleModel, 0, len(roles))
	roleMap := make(map[int64]*v1.RoleModel)
	for _, role := range roles {
		roleModel := c.roleModelToV1(role)
		models = append(models, roleModel)
		roleMap[role.Id] = roleModel
	}

	// 排序和构建角色树
	c.sortRoleModels(models)
	tree := c.buildRoleTree(models, roleMap)

	// 构建响应
	res = &v1.GetListRes{
		List: models,
		Tree: tree,
	}

	return
}

func (c *ControllerV1) sortRoleModels(models []*v1.RoleModel) {
	slices.SortFunc(models, func(a, b *v1.RoleModel) int {
		return cmp.Compare(a.Sort, b.Sort)
	})
}

// 角色树构建
func (c *ControllerV1) buildRoleTree(roles []*v1.RoleModel, roleMap map[int64]*v1.RoleModel) []*v1.RoleModel {
	tree := make([]*v1.RoleModel, 0)
	for _, node := range roles {
		parentId := node.ParentId
		if parent, exists := roleMap[parentId]; exists {
			parent.Children = append(parent.Children, node)
		} else if parentId == 0 {
			tree = append(tree, node)
		}
	}
	return tree
}
