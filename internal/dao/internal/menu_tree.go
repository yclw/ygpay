// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MenuTreeDao is the data access object for the table t_menu_tree.
type MenuTreeDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MenuTreeColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MenuTreeColumns defines and stores column names for the table t_menu_tree.
type MenuTreeColumns struct {
	Id  string // 菜单ID
	Pid string // 父菜单ID
}

// menuTreeColumns holds the columns for the table t_menu_tree.
var menuTreeColumns = MenuTreeColumns{
	Id:  "id",
	Pid: "pid",
}

// NewMenuTreeDao creates and returns a new DAO object for table data access.
func NewMenuTreeDao(handlers ...gdb.ModelHandler) *MenuTreeDao {
	return &MenuTreeDao{
		group:    "default",
		table:    "t_menu_tree",
		columns:  menuTreeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MenuTreeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MenuTreeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MenuTreeDao) Columns() MenuTreeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MenuTreeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MenuTreeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MenuTreeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
