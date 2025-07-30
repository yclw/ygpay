package role

import (
	"context"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/internal/logic/role"
	"yclw/ygpay/util/tree"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	_, idTree, err := c.RoleService.GetAllList(ctx)
	if err != nil {
		return
	}
	res = &v1.GetListRes{}
	res.Tree = c.roleModelToV1Trees(idTree)
	return
}

// roleModelToV1Trees 构建RoleTreeModel列表
func (c *ControllerV1) roleModelToV1Trees(idTree *tree.IdTree) []*v1.RoleTreeModel {
	res := make([]*v1.RoleTreeModel, 0, len(idTree.Root.Children))
	for _, node := range idTree.Root.Children {
		res = append(res, c.roleModelToV1Tree(node))
	}
	return res
}

// roleModelToV1Tree 递归构建RoleTreeModel
func (c *ControllerV1) roleModelToV1Tree(node *tree.TreeNode) *v1.RoleTreeModel {
	// 构建RoleTreeModel
	roleModel := node.Data.(*role.RoleModel)
	roleTreeModel := &v1.RoleTreeModel{
		RoleModel: &v1.RoleModel{
			Id:   roleModel.RoleInfo.Id,
			Name: roleModel.RoleInfo.Name,
			Key:  roleModel.RoleInfo.Key,
		},
	}
	// 递归构建子节点
	for _, child := range node.Children {
		roleTreeModel.Children = append(roleTreeModel.Children, c.roleModelToV1Tree(child))
	}
	return roleTreeModel
}
