package dao

import (
	"github.com/gogf/gf/v2/database/gdb"
)

// QueryOption 查询选项函数类型
type QueryOption func(*gdb.Model) *gdb.Model

// WithStatus 添加状态筛选
func WithStatus(status int) QueryOption {
	return func(model *gdb.Model) *gdb.Model {
		return model.Where("status", status)
	}
}

// Where 添加条件筛选
func Where(where interface{}, args ...interface{}) QueryOption {
	return func(model *gdb.Model) *gdb.Model {
		return model.Where(where, args)
	}
}

// WhereIn 添加条件筛选
func WhereIn(column string, in interface{}) QueryOption {
	return func(model *gdb.Model) *gdb.Model {
		return model.WhereIn(column, in)
	}
}

// WhereLike 添加条件筛选
func WhereLike(column string, like string) QueryOption {
	return func(model *gdb.Model) *gdb.Model {
		return model.WhereLike(column, like)
	}
}

// WhereBetween 添加条件筛选
func WhereBetween(column string, min interface{}, max interface{}) QueryOption {
	return func(model *gdb.Model) *gdb.Model {
		return model.WhereBetween(column, min, max)
	}
}

// OrderAsc 按照指定字段递增排序
func OrderAsc(column string) QueryOption {
	return func(model *gdb.Model) *gdb.Model {
		return model.OrderAsc(column)
	}
}

// OrderDesc 按照指定字段递减排序
func OrderDesc(column string) QueryOption {
	return func(model *gdb.Model) *gdb.Model {
		return model.OrderDesc(column)
	}
}

// applyOptions 应用查询选项
func applyOptions(model *gdb.Model, options ...QueryOption) *gdb.Model {
	for _, option := range options {
		if option != nil {
			model = option(model)
		}
	}
	return model
}
