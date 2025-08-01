// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"context"
	"yclw/ygpay/internal/dao/internal"
	"yclw/ygpay/internal/model/do"
	"yclw/ygpay/internal/model/entity"
	"yclw/ygpay/util/tree"
)

// menuTreeDao is the data access object for the table t_menu_tree.
// You can define custom methods on it to extend its functionality as needed.
type menuTreeDao struct {
	*internal.MenuTreeDao
}

var (
	// MenuTree is a globally accessible object for table t_menu_tree operations.
	MenuTree = menuTreeDao{internal.NewMenuTreeDao()}
)

// Add your custom methods and functionality below.

// BuildTreeByMenuIds 根据菜单ID构建树
func (d *menuTreeDao) BuildTreeByMenuIds(ctx context.Context, menuIds []int64) (res *tree.IdTree, err error) {
	T := []tree.T{}
	cols := d.Columns()
	err = d.Ctx(ctx).WhereIn(cols.Id, menuIds).Scan(&T)
	if err != nil {
		return
	}
	T = append(T, tree.T{Id: 0, Pid: -1})
	res, err = tree.BuildTree(T, 0)
	return
}

// FindAll 获取所有菜单树
func (d *menuTreeDao) FindAll(ctx context.Context) (res []tree.T, err error) {
	err = d.Ctx(ctx).Scan(&res)
	return
}

// FindByMenuIds 根据菜单ID获取菜单树
func (d *menuTreeDao) FindByMenuIds(ctx context.Context, menuIds []int64) (res []tree.T, err error) {
	err = d.Ctx(ctx).WhereIn(d.Columns().Id, menuIds).Scan(&res)
	return
}

// FindPidById 根据菜单ID获取父菜单ID
func (d *menuTreeDao) FindPidById(ctx context.Context, id int64) (res int64, err error) {
	model := entity.MenuTree{}
	cols := d.Columns()
	err = d.Ctx(ctx).Where(cols.Id, id).Scan(&model)
	if err != nil {
		return
	}
	res = model.Pid
	return
}

// Create 创建菜单树
func (d *menuTreeDao) Create(ctx context.Context, req *do.MenuTree) (id int64, err error) {
	cols := d.Columns()
	mod, err := d.Ctx(ctx).Fields(
		cols.Id,
		cols.Pid,
	).Data(req).OmitNil().Insert()
	if err != nil {
		return
	}
	id, err = mod.LastInsertId()
	return
}

// Update 更新菜单树
func (d *menuTreeDao) Update(ctx context.Context, req *do.MenuTree) (err error) {
	cols := d.Columns()
	_, err = d.Ctx(ctx).Where(cols.Id, req.Id).Fields(
		cols.Pid,
	).Data(req).OmitNil().Update()
	return
}

// Delete 删除菜单树
func (d *menuTreeDao) Delete(ctx context.Context, id int64) (err error) {
	cols := d.Columns()
	_, err = d.Ctx(ctx).Where(cols.Id, id).Delete()
	return
}
