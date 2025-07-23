// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttachmentDao is the data access object for the table t_attachment.
type AttachmentDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AttachmentColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AttachmentColumns defines and stores column names for the table t_attachment.
type AttachmentColumns struct {
	Id        string // 文件ID
	AppId     string // 应用ID
	MemberId  string // 管理员ID
	CateId    string // 上传分类
	Drive     string // 上传驱动
	Name      string // 文件原始名
	Kind      string // 上传类型
	MimeType  string // 扩展类型
	NaiveType string // NaiveUI类型
	Path      string // 本地路径
	FileUrl   string // url
	Size      string // 文件大小
	Ext       string // 扩展名
	Md5       string // md5校验码
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

// attachmentColumns holds the columns for the table t_attachment.
var attachmentColumns = AttachmentColumns{
	Id:        "id",
	AppId:     "app_id",
	MemberId:  "member_id",
	CateId:    "cate_id",
	Drive:     "drive",
	Name:      "name",
	Kind:      "kind",
	MimeType:  "mime_type",
	NaiveType: "naive_type",
	Path:      "path",
	FileUrl:   "file_url",
	Size:      "size",
	Ext:       "ext",
	Md5:       "md5",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAttachmentDao creates and returns a new DAO object for table data access.
func NewAttachmentDao(handlers ...gdb.ModelHandler) *AttachmentDao {
	return &AttachmentDao{
		group:    "default",
		table:    "t_attachment",
		columns:  attachmentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttachmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttachmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttachmentDao) Columns() AttachmentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttachmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttachmentDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AttachmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
