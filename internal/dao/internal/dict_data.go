// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DictDataDao is the data access object for the table t_dict_data.
type DictDataDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DictDataColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DictDataColumns defines and stores column names for the table t_dict_data.
type DictDataColumns struct {
	Id        string // 字典数据ID
	Label     string // 字典标签
	Value     string // 字典键值
	ValueType string // 键值数据类型：string,int,uint,bool,datetime,date
	Type      string // 字典类型
	ListClass string // 表格回显样式
	IsDefault string // 是否为系统默认
	Sort      string // 字典排序
	Remark    string // 备注
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// dictDataColumns holds the columns for the table t_dict_data.
var dictDataColumns = DictDataColumns{
	Id:        "id",
	Label:     "label",
	Value:     "value",
	ValueType: "value_type",
	Type:      "type",
	ListClass: "list_class",
	IsDefault: "is_default",
	Sort:      "sort",
	Remark:    "remark",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewDictDataDao creates and returns a new DAO object for table data access.
func NewDictDataDao(handlers ...gdb.ModelHandler) *DictDataDao {
	return &DictDataDao{
		group:    "default",
		table:    "t_dict_data",
		columns:  dictDataColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DictDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DictDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DictDataDao) Columns() DictDataColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DictDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DictDataDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DictDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
