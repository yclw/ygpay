package role

import (
	"context"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/util/tree"
)

var RoleService = NewRole()

type Role struct {
}

func NewRole() *Role {
	return &Role{}
}

// GetOne 获取单个角色信息
func (r *Role) GetOne(ctx context.Context, id int64) (res *RoleModel, err error) {
	res = &RoleModel{}

	// 获取角色信息
	role, err := dao.RoleInfo.FindByID(ctx, id)
	if err != nil {
		return
	}
	res.RoleInfo = role

	// 获得pid
	pid, err := dao.RoleTree.FindPidById(ctx, id)
	if err != nil {
		return
	}
	res.Pid = pid

	return
}

// GetAllList 获取所有角色列表
func (r *Role) GetAllList(ctx context.Context) (res []*RoleModel, idTree *tree.IdTree, err error) {
	// 获取所有角色信息
	roles, err := dao.RoleInfo.FindAll(ctx)
	if err != nil {
		return
	}

	// 获取所有角色id树关系
	Ts, err := dao.RoleTree.FindAll(ctx)
	if err != nil {
		return
	}

	// 构建角色树
	Ts = append(Ts, tree.T{Id: 0, Pid: -1})
	idTree, err = tree.BuildTree(Ts, 0)
	if err != nil {
		return
	}

	// 转换为RoleModel
	res = make([]*RoleModel, 0, len(roles))
	for _, role := range roles {
		res = append(res, &RoleModel{RoleInfo: role, TreeNode: idTree.NodeMap[role.Id]})
		idTree.NodeMap[role.Id].Data = role
	}

	return
}

// Create 创建角色
func (r *Role) Create(ctx context.Context, req *RoleCreateModel) (err error) {
	// 创建角色
	id, err := dao.RoleInfo.Create(ctx, req.RoleInfo)
	if err != nil {
		return
	}

	req.RoleTree.Id = id

	// 创建角色树
	_, err = dao.RoleTree.Create(ctx, req.RoleTree)
	if err != nil {
		return
	}

	return
}

// Update 更新角色
func (r *Role) Update(ctx context.Context, req *RoleUpdateModel) (err error) {
	// 更新角色
	err = dao.RoleInfo.Update(ctx, req.RoleInfo)
	if err != nil {
		return
	}

	// 更新角色树
	err = dao.RoleTree.Update(ctx, req.RoleTree)
	if err != nil {
		return
	}

	return
}

// Delete 删除角色
func (r *Role) Delete(ctx context.Context, id int64) (err error) {
	// 删除角色
	err = dao.RoleInfo.Delete(ctx, id)
	if err != nil {
		return
	}

	// 删除角色树
	err = dao.RoleTree.Delete(ctx, id)
	if err != nil {
		return
	}

	return
}
