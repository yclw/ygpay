package menu

import (
	"yclw/ygpay/internal/model/do"
	"yclw/ygpay/internal/model/entity"
	"yclw/ygpay/util/tree"
)

type RoleMenuModel struct {
	*entity.MenuInfo
	*tree.TreeNode
}

type MenuModel struct {
	*entity.MenuInfo
	*tree.TreeNode
}

type MenuCreateModel struct {
	Id int64
	*do.MenuInfo
	*do.MenuTree
}

type MenuUpdateModel struct {
	Id int64
	*do.MenuInfo
	*do.MenuTree
}

// MenuListFilter 菜单列表筛选参数
type MenuListFilter struct {
	Status *int `json:"status"` // 状态筛选
}
