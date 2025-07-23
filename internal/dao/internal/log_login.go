// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LogLoginDao is the data access object for the table t_log_login.
type LogLoginDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  LogLoginColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// LogLoginColumns defines and stores column names for the table t_log_login.
type LogLoginColumns struct {
	Id            string // 登录日志ID
	MemberId      string // 用户ID
	Username      string // 登录账号
	IpAddress     string // IP地址
	UserAgent     string // 用户代理
	LoginLocation string // 登录地点
	Browser       string // 浏览器
	Os            string // 操作系统
	LoginStatus   string // 登录状态: 0失败 1成功
	LoginMessage  string // 登录信息
	LoginTime     string // 登录时间
}

// logLoginColumns holds the columns for the table t_log_login.
var logLoginColumns = LogLoginColumns{
	Id:            "id",
	MemberId:      "member_id",
	Username:      "username",
	IpAddress:     "ip_address",
	UserAgent:     "user_agent",
	LoginLocation: "login_location",
	Browser:       "browser",
	Os:            "os",
	LoginStatus:   "login_status",
	LoginMessage:  "login_message",
	LoginTime:     "login_time",
}

// NewLogLoginDao creates and returns a new DAO object for table data access.
func NewLogLoginDao(handlers ...gdb.ModelHandler) *LogLoginDao {
	return &LogLoginDao{
		group:    "default",
		table:    "t_log_login",
		columns:  logLoginColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LogLoginDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LogLoginDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LogLoginDao) Columns() LogLoginColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LogLoginDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LogLoginDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *LogLoginDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
