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
	Id             string // 菜单ID
	Pid            string // 父菜单ID
	Level          string // 关系树等级
	Tree           string // 关系树
	Title          string // 菜单名称
	Name           string // 名称编码
	Path           string // 路由地址
	Icon           string // 菜单图标
	Type           string // 菜单类型（1目录 2菜单 3按钮）
	Redirect       string // 重定向地址
	Permissions    string // 菜单包含权限集合
	PermissionName string // 权限名称
	Component      string // 组件路径
	AlwaysShow     string // 取消自动计算根路由模式
	ActiveMenu     string // 高亮菜单编码
	IsRoot         string // 是否跟路由
	IsFrame        string // 是否内嵌
	FrameSrc       string // 内联外部地址
	KeepAlive      string // 缓存该路由
	Hidden         string // 是否隐藏
	Affix          string // 是否固定
	Sort           string // 排序
	Remark         string // 备注
	Status         string // 菜单状态
	UpdatedAt      string // 更新时间
	CreatedAt      string // 创建时间
}

// menuInfoColumns holds the columns for the table t_menu_info.
var menuInfoColumns = MenuInfoColumns{
	Id:             "id",
	Pid:            "pid",
	Level:          "level",
	Tree:           "tree",
	Title:          "title",
	Name:           "name",
	Path:           "path",
	Icon:           "icon",
	Type:           "type",
	Redirect:       "redirect",
	Permissions:    "permissions",
	PermissionName: "permission_name",
	Component:      "component",
	AlwaysShow:     "always_show",
	ActiveMenu:     "active_menu",
	IsRoot:         "is_root",
	IsFrame:        "is_frame",
	FrameSrc:       "frame_src",
	KeepAlive:      "keep_alive",
	Hidden:         "hidden",
	Affix:          "affix",
	Sort:           "sort",
	Remark:         "remark",
	Status:         "status",
	UpdatedAt:      "updated_at",
	CreatedAt:      "created_at",
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
