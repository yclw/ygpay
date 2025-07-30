// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MemberInfoDao is the data access object for the table t_member_info.
type MemberInfoDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MemberInfoColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MemberInfoColumns defines and stores column names for the table t_member_info.
type MemberInfoColumns struct {
	Id           string // 用户ID
	Uid          string // 用户UID
	Username     string // 帐号
	PasswordHash string // 密码哈希
	Avatar       string // 头像
	Sex          string // 性别: 1男 2女 3未知
	Email        string // 邮箱
	Mobile       string // 手机号码
	Address      string // 联系地址
	LastActiveAt string // 最后活跃时间
	Remark       string // 备注
	Sort         string // 排序
	Status       string // 状态: 0禁用 1启用
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	Nickname     string // 昵称
}

// memberInfoColumns holds the columns for the table t_member_info.
var memberInfoColumns = MemberInfoColumns{
	Id:           "id",
	Uid:          "uid",
	Username:     "username",
	PasswordHash: "password_hash",
	Avatar:       "avatar",
	Sex:          "sex",
	Email:        "email",
	Mobile:       "mobile",
	Address:      "address",
	LastActiveAt: "last_active_at",
	Remark:       "remark",
	Sort:         "sort",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	Nickname:     "nickname",
}

// NewMemberInfoDao creates and returns a new DAO object for table data access.
func NewMemberInfoDao(handlers ...gdb.ModelHandler) *MemberInfoDao {
	return &MemberInfoDao{
		group:    "default",
		table:    "t_member_info",
		columns:  memberInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MemberInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MemberInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MemberInfoDao) Columns() MemberInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MemberInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MemberInfoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MemberInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
