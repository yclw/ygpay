// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ApiInfoDao is the data access object for the table t_api_info.
type ApiInfoDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ApiInfoColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ApiInfoColumns defines and stores column names for the table t_api_info.
type ApiInfoColumns struct {
	Id          string // API ID
	Name        string // API名称
	Path        string // API路径
	Method      string // API方法
	GroupName   string // API分组
	Description string // API描述
	NeedAuth    string // 是否需要认证: 0否 1是
	RateLimit   string // 限流次数/分钟
	Sort        string // 排序
	Status      string // 状态: 0禁用 1启用
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// apiInfoColumns holds the columns for the table t_api_info.
var apiInfoColumns = ApiInfoColumns{
	Id:          "id",
	Name:        "name",
	Path:        "path",
	Method:      "method",
	GroupName:   "group_name",
	Description: "description",
	NeedAuth:    "need_auth",
	RateLimit:   "rate_limit",
	Sort:        "sort",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewApiInfoDao creates and returns a new DAO object for table data access.
func NewApiInfoDao(handlers ...gdb.ModelHandler) *ApiInfoDao {
	return &ApiInfoDao{
		group:    "default",
		table:    "t_api_info",
		columns:  apiInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ApiInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ApiInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ApiInfoDao) Columns() ApiInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ApiInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ApiInfoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ApiInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
