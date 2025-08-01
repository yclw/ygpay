package role

import (
	"yclw/ygpay/internal/model/do"
	"yclw/ygpay/internal/model/entity"
	"yclw/ygpay/util/tree"

	"github.com/gogf/gf/v2/os/gtime"
)

// RoleModel 角色模型
type RoleModel struct {
	*entity.RoleInfo
	*entity.RoleMenu
	*entity.RoleApi
	*tree.TreeNode
	ParentName string
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

// RoleListFilter 角色列表筛选参数
type RoleListFilter struct {
	Name      string      `json:"name"`      // 角色名称
	Key       string      `json:"key"`       // 角色权限字符串
	Status    *int        `json:"status"`    // 状态筛选
	StartDate *gtime.Time `json:"startDate"` // 开始日期
	EndDate   *gtime.Time `json:"endDate"`   // 结束日期
	SortField string      `json:"sortField"` // 排序字段
	SortDesc  bool        `json:"sortDesc"`  // 是否降序
}
