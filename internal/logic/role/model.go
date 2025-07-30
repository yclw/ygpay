package role

import (
	"yclw/ygpay/internal/model/do"
	"yclw/ygpay/internal/model/entity"
	"yclw/ygpay/util/tree"
)

// RoleModel 角色模型
type RoleModel struct {
	*entity.RoleInfo
	*entity.RoleMenu
	*entity.RoleApi
	*tree.TreeNode
}

// RoleUpdateModel 角色更新模型
type RoleUpdateModel struct {
	Id int64 `json:"id"                 dc:"角色ID"`
	*do.RoleInfo
	*do.RoleMenu
	*do.RoleApi
	*do.RoleTree
}

// RoleCreateModel 角色创建模型
type RoleCreateModel struct {
	Id int64 `json:"id"                 dc:"角色ID"`
	*do.RoleInfo
	*do.RoleMenu
	*do.RoleApi
	*do.RoleTree
}
