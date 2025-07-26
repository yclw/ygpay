package menu

import (
	"context"
	"yclw/ygpay/internal/dao"
)

var MenuService = NewMenu()

type Menu struct {
}

func NewMenu() *Menu {
	return &Menu{}
}

// GetOne 获取单个菜单信息
func (m *Menu) GetOne(ctx context.Context, id int64) (res *MenuModel, err error) {
	res = &MenuModel{}

	// 获取菜单信息
	menu, err := dao.MenuInfo.FindByID(ctx, id)
	if err != nil {
		return
	}
	res.MenuInfo = menu
	return
}

// GetAllList 获取所有菜单列表
func (m *Menu) GetAllList(ctx context.Context) (res []*MenuModel, err error) {
	// 获取所有菜单信息
	menus, err := dao.MenuInfo.FindAll(ctx)
	if err != nil {
		return
	}

	res = make([]*MenuModel, 0, len(menus))
	for _, menu := range menus {
		res = append(res, &MenuModel{MenuInfo: menu})
	}

	// 获取所有菜单id树
	return
}

// Create 创建菜单
func (m *Menu) Create(ctx context.Context, req *MenuCreateModel) (err error) {
	return
}

// Update 更新菜单
func (m *Menu) Update(ctx context.Context, req *MenuUpdateModel) (err error) {
	return
}

// Delete 删除菜单
func (m *Menu) Delete(ctx context.Context, id int64) (err error) {
	return
}
