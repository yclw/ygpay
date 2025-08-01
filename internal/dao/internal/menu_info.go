// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MenuInfoDao is the data access object for the table t_menu_info.
type MenuInfoDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MenuInfoColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MenuInfoColumns defines and stores column names for the table t_menu_info.
type MenuInfoColumns struct {
	Id         string // 菜单ID
	Type       string // 菜单类型: 0目录 1菜单 2外链
	Name       string // 菜单名称
	Path       string // 菜单路径
	Title      string // 菜单标题
	Icon       string // 菜单图标
	Sort       string // 排序
	ShowParent string // 是否显示父菜单: 0是 1否
	ShowLink   string // 是否显示该菜单: 0是 1否
	KeepAlive  string // 是否缓存: 0是 1否
	Redirect   string // 重定向
	Component  string // 组件路径
	FrameSrc   string // 内嵌地址
	Url        string // 外部链接
	Status     string // 状态: 0禁用 1启用
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// menuInfoColumns holds the columns for the table t_menu_info.
var menuInfoColumns = MenuInfoColumns{
	Id:         "id",
	Type:       "type",
	Name:       "name",
	Path:       "path",
	Title:      "title",
	Icon:       "icon",
	Sort:       "sort",
	ShowParent: "showParent",
	ShowLink:   "showLink",
	KeepAlive:  "keepAlive",
	Redirect:   "redirect",
	Component:  "component",
	FrameSrc:   "frameSrc",
	Url:        "url",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewMenuInfoDao creates and returns a new DAO object for table data access.
func NewMenuInfoDao(handlers ...gdb.ModelHandler) *MenuInfoDao {
	return &MenuInfoDao{
		group:    "default",
		table:    "t_menu_info",
		columns:  menuInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MenuInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MenuInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MenuInfoDao) Columns() MenuInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MenuInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MenuInfoDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *MenuInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
