package role

import (
	"context"
	"fmt"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/model/entity"
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
	res = &RoleModel{
		RoleInfo: &entity.RoleInfo{},
		TreeNode: &tree.TreeNode{},
	}

	// 获取角色信息
	role, err := dao.RoleInfo.FindByID(ctx, id)
	if err != nil {
		fmt.Println("FindByID", err)
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

// GetListWithFilter 获取角色列表（带筛选）
func (r *Role) GetListWithFilter(ctx context.Context, page, size int, filter *RoleListFilter) (res []*RoleModel, total int, err error) {
	// 构建查询选项
	options := r.buildQueryOptions(filter)

	// 获取筛选后的角色信息
	roles, total, err := dao.RoleInfo.FindWithPageAndOptions(ctx, page, size, options...)
	if err != nil {
		return
	}

	// 获取所有角色id树关系
	Ts, err := dao.RoleTree.FindAll(ctx)
	if err != nil {
		return
	}
	pidMap := make(map[int64]int64)
	for _, t := range Ts {
		pidMap[t.Id] = t.Pid
	}

	// 创建角色信息映射表，用于快速查找父角色名称
	allRoles, err := dao.RoleInfo.FindAll(ctx)
	if err != nil {
		return
	}
	roleMap := make(map[int64]*entity.RoleInfo)
	for _, role := range allRoles {
		roleMap[role.Id] = role
	}

	// 转换为RoleModel
	res = make([]*RoleModel, 0, len(roles))
	for _, role := range roles {
		parentName := ""
		if parentRole, exists := roleMap[pidMap[role.Id]]; exists {
			parentName = parentRole.Name
		}
		roleModel := &RoleModel{
			RoleInfo:   role,
			TreeNode:   &tree.TreeNode{Id: role.Id, Pid: pidMap[role.Id]},
			ParentName: parentName,
		}
		res = append(res, roleModel)
	}

	return
}

// buildQueryOptions 构建查询选项
func (r *Role) buildQueryOptions(filter *RoleListFilter) []dao.QueryOption {
	if filter == nil {
		return nil
	}

	var options []dao.QueryOption
	cols := dao.RoleInfo.Columns()

	// 角色名称筛选
	if filter.Name != "" {
		options = append(options, dao.WhereLike(cols.Name, "%"+filter.Name+"%"))
	}

	// 角色权限字符串筛选
	if filter.Key != "" {
		options = append(options, dao.WhereLike(cols.Key, "%"+filter.Key+"%"))
	}

	// 状态筛选
	if filter.Status != nil {
		options = append(options, dao.Where(cols.Status, *filter.Status))
	}

	// 日期范围筛选
	if filter.StartDate != nil || filter.EndDate != nil {
		options = append(options, dao.WhereBetween(cols.CreatedAt, filter.StartDate, filter.EndDate))
	}

	// 自定义排序
	if filter.SortField != "" {
		if filter.SortDesc {
			options = append(options, dao.OrderDesc(filter.SortField))
		} else {
			options = append(options, dao.OrderAsc(filter.SortField))
		}
	}

	return options
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
