package role

import (
	"context"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/model/entity"
)

var RoleService = NewRole()

type Role struct {
}

func NewRole() *Role {
	return &Role{}
}

// GetOne 获取单个角色信息
func (r *Role) GetOne(ctx context.Context, id int64) (res *RoleModel, err error) {
	// 创建角色模型
	res = &RoleModel{
		RoleInfo: &entity.RoleInfo{},
	}

	// 获取角色信息
	res.RoleInfo, err = dao.RoleInfo.FindByID(ctx, id)
	if err != nil {
		return
	}

	// 获取父角色名称
	pRole, err := dao.RoleInfo.FindByID(ctx, res.RoleInfo.Pid)
	if err == nil && pRole != nil {
		res.ParentName = pRole.Name
	}

	return
}

// GetAllList 获取所有角色列表
func (r *Role) GetAllList(ctx context.Context) (res []*RoleModel, err error) {
	// 获取所有角色信息
	roles, err := dao.RoleInfo.FindAll(ctx)
	if err != nil {
		return
	}

	// 创建角色信息映射表
	roleMap := make(map[int64]*entity.RoleInfo)
	for _, role := range roles {
		roleMap[role.Id] = role
	}

	// 转换为RoleModel
	res = make([]*RoleModel, 0, len(roles))
	for _, role := range roles {
		// 创建角色模型
		roleModel := &RoleModel{RoleInfo: role}

		// 获取父角色名称
		if parentRole, exists := roleMap[role.Pid]; exists {
			roleModel.ParentName = parentRole.Name
		}

		// 添加到结果
		res = append(res, roleModel)
	}

	return
}

// Create 创建角色
func (r *Role) Create(ctx context.Context, req *RoleCreateModel) (err error) {
	// 创建角色
	_, err = dao.RoleInfo.Create(ctx, req.RoleInfo)
	return
}

// Update 更新角色
func (r *Role) Update(ctx context.Context, req *RoleUpdateModel) (err error) {
	// 更新角色
	err = dao.RoleInfo.Update(ctx, req.RoleInfo)
	return
}

// Delete 删除角色
func (r *Role) Delete(ctx context.Context, id int64) (err error) {
	// 删除角色
	err = dao.RoleInfo.Delete(ctx, id)
	return
}
