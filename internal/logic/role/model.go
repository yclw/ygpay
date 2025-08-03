package role

import (
	"yclw/ygpay/internal/model/do"
	"yclw/ygpay/internal/model/entity"
)

// RoleModel 角色模型
type RoleModel struct {
	*entity.RoleInfo        // 角色信息
	ParentName       string // 父角色名称
}

// RoleUpdateModel 角色更新模型
type RoleUpdateModel struct {
	*do.RoleInfo
}

// RoleCreateModel 角色创建模型
type RoleCreateModel struct {
	*do.RoleInfo
}
