package menu

import (
	"yclw/ygpay/internal/model/do"
	"yclw/ygpay/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

type RoleMenuModel struct {
	*entity.MenuInfo
}

type MenuModel struct {
	*entity.MenuInfo
	ParentTitle string
}

type MenuCreateModel struct {
	*do.MenuInfo
}

type MenuUpdateModel struct {
	*do.MenuInfo
}

// MenuListFilter 菜单列表筛选参数
type MenuListFilter struct {
	Status    *int        `json:"status"`    // 状态筛选
	Type      *int        `json:"type"`      // 菜单类型: 0目录 1菜单 2外链
	Name      string      `json:"name"`      // 菜单名称
	Path      string      `json:"path"`      // 菜单路径
	Title     string      `json:"title"`     // 菜单标题
	StartDate *gtime.Time `json:"startDate"` // 开始日期
	EndDate   *gtime.Time `json:"endDate"`   // 结束日期
	SortField string      `json:"sortField"` // 排序字段
	SortDesc  bool        `json:"sortDesc"`  // 是否降序
}

const (
	MenuTypeDir  = 0
	MenuTypeMenu = 1
	MenuTypeLink = 2
)
