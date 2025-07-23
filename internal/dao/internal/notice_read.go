// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NoticeReadDao is the data access object for the table t_notice_read.
type NoticeReadDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  NoticeReadColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// NoticeReadColumns defines and stores column names for the table t_notice_read.
type NoticeReadColumns struct {
	Id        string // 记录ID
	NoticeId  string // 公告ID
	MemberId  string // 会员ID
	Clicks    string // 已读次数
	UpdatedAt string // 更新时间
	CreatedAt string // 阅读时间
}

// noticeReadColumns holds the columns for the table t_notice_read.
var noticeReadColumns = NoticeReadColumns{
	Id:        "id",
	NoticeId:  "notice_id",
	MemberId:  "member_id",
	Clicks:    "clicks",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
}

// NewNoticeReadDao creates and returns a new DAO object for table data access.
func NewNoticeReadDao(handlers ...gdb.ModelHandler) *NoticeReadDao {
	return &NoticeReadDao{
		group:    "default",
		table:    "t_notice_read",
		columns:  noticeReadColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *NoticeReadDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *NoticeReadDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *NoticeReadDao) Columns() NoticeReadColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *NoticeReadDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *NoticeReadDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *NoticeReadDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
