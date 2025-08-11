package role

import (
	"context"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/model/entity"

	"github.com/google/uuid"
)

var RoleService = NewRole()

type Role struct {
}

func NewRole() *Role {
	return &Role{}
}

// GetOne 获取单个角色信息
func (r *Role) GetOne(ctx context.Context, roleUid string) (res *RoleModel, err error) {
	// 创建角色模型
	res = &RoleModel{
		RoleInfo: &entity.RoleInfo{},
	}

	// 获取角色信息
	res.RoleInfo, err = dao.RoleInfo.FindByRoleUid(ctx, roleUid)
	if err != nil {
		return
	}

	// 获取父角色名称和UID
	pRole, err := dao.RoleInfo.FindByID(ctx, res.RoleInfo.Pid)
	if err == nil && pRole != nil {
		res.ParentName = pRole.Name
		res.ParentUid = pRole.RoleUid
	}

	return
}

// GetRoleIdByUid 根据roleUid获取roleId
func (r *Role) GetRoleIdByUid(ctx context.Context, roleUid string) (res int64, err error) {
	res, err = dao.RoleInfo.FindIdByRoleUid(ctx, roleUid)
	return
}

// GetOneById 根据ID获取单个角色信息
func (r *Role) GetOneById(ctx context.Context, id int64) (res *RoleModel, err error) {
	// 创建角色模型
	res = &RoleModel{
		RoleInfo: &entity.RoleInfo{},
	}

	// 获取角色信息
	res.RoleInfo, err = dao.RoleInfo.FindByID(ctx, id)
	if err != nil {
		return
	}

	// 获取父角色名称和UID
	pRole, err := dao.RoleInfo.FindByID(ctx, res.RoleInfo.Pid)
	if err == nil && pRole != nil {
		res.ParentName = pRole.Name
		res.ParentUid = pRole.RoleUid
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

		// 获取父角色名称和UID
		if parentRole, exists := roleMap[role.Pid]; exists {
			roleModel.ParentName = parentRole.Name
			roleModel.ParentUid = parentRole.RoleUid
		}

		// 添加到结果
		res = append(res, roleModel)
	}

	return
}

// Create 创建角色
func (r *Role) Create(ctx context.Context, req *RoleCreateModel) (roleUid string, err error) {
	// 生成RoleUid
	roleUid = uuid.New().String()
	req.RoleInfo.RoleUid = roleUid

	// 处理父角色ID
	if req.ParentUid != "" {
		req.RoleInfo.Pid, err = dao.RoleInfo.FindIdByRoleUid(ctx, req.ParentUid)
		if err != nil {
			return
		}
	} else {
		req.RoleInfo.Pid = 0
	}

	// 创建角色
	_, err = dao.RoleInfo.Create(ctx, req.RoleInfo)
	return
}

// Update 更新角色
func (r *Role) Update(ctx context.Context, req *RoleUpdateModel) (err error) {
	// 处理父角色ID
	if req.ParentUid != "" {
		req.RoleInfo.Pid, err = dao.RoleInfo.FindIdByRoleUid(ctx, req.ParentUid)
		if err != nil {
			return
		}
	} else {
		req.RoleInfo.Pid = 0
	}

	// 更新角色
	err = dao.RoleInfo.UpdateByRoleUid(ctx, req.RoleInfo)
	return
}

// Delete 删除角色
func (r *Role) Delete(ctx context.Context, roleUid string) (err error) {
	// 删除角色
	err = dao.RoleInfo.DeleteByRoleUid(ctx, roleUid)
	return
}
