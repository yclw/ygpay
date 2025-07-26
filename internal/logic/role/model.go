package role

import (
	"yclw/ygpay/internal/model/do"
	"yclw/ygpay/internal/model/entity"
)

// RoleModel 角色模型
type RoleModel struct {
	*entity.RoleInfo
	*entity.RoleMenu
	*entity.RoleApi
}

// RoleUpdateModel 角色更新模型
type RoleUpdateModel struct {
	Id int64 `json:"id"                 dc:"角色ID"`
	*do.RoleInfo
	*do.RoleMenu
	*do.RoleApi
}

// RoleCreateModel 角色创建模型
type RoleCreateModel struct {
	Id int64 `json:"id"                 dc:"角色ID"`
	*do.RoleInfo
	*do.RoleMenu
	*do.RoleApi
}
