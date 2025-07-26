// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoleInfoDao is the data access object for the table t_role_info.
type RoleInfoDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  RoleInfoColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// RoleInfoColumns defines and stores column names for the table t_role_info.
type RoleInfoColumns struct {
	Id        string // 角色ID
	Name      string // 角色名称
	Key       string // 角色权限字符串
	Remark    string // 备注
	Sort      string // 排序
	Status    string // 状态: 0禁用 1启用
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// roleInfoColumns holds the columns for the table t_role_info.
var roleInfoColumns = RoleInfoColumns{
	Id:        "id",
	Name:      "name",
	Key:       "key",
	Remark:    "remark",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewRoleInfoDao creates and returns a new DAO object for table data access.
func NewRoleInfoDao(handlers ...gdb.ModelHandler) *RoleInfoDao {
	return &RoleInfoDao{
		group:    "default",
		table:    "t_role_info",
		columns:  roleInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RoleInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RoleInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RoleInfoDao) Columns() RoleInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RoleInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RoleInfoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RoleInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
