package role

import "context"

var RoleService = NewRole()

type Role struct {
}

func NewRole() *Role {
	return &Role{}
}

// GetOne 获取单个角色信息
func (r *Role) GetOne(ctx context.Context, id int64) (res *RoleModel, err error) {
	return
}

// GetAllList 获取所有角色列表
func (r *Role) GetAllList(ctx context.Context) (res []*RoleModel, err error) {
	return
}

// GetSubList 获取子角色列表
func (r *Role) GetSubList(ctx context.Context, id int64) (res []*RoleModel, err error) {
	return
}

// Create 创建角色
func (r *Role) Create(ctx context.Context, req *RoleCreateModel) (err error) {
	return
}

// Update 更新角色
func (r *Role) Update(ctx context.Context, req *RoleUpdateModel) (err error) {
	return
}

// Delete 删除角色
func (r *Role) Delete(ctx context.Context, id int64) (err error) {
	return
}
