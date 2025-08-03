package menu

import (
	"context"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/model/entity"
)

var MenuService = NewMenu()

type Menu struct {
}

func NewMenu() *Menu {
	return &Menu{}
}

// GetOne 获取单个菜单信息
func (m *Menu) GetOne(ctx context.Context, id int64) (res *MenuModel, err error) {
	// 创建菜单模型
	res = &MenuModel{
		MenuInfo: &entity.MenuInfo{},
	}

	// 获取菜单信息
	res.MenuInfo, err = dao.MenuInfo.FindByID(ctx, id)
	if err != nil {
		return
	}

	// 获取父菜单信息
	pMenu, err := dao.MenuInfo.FindByID(ctx, res.MenuInfo.Pid)
	if err == nil && pMenu != nil {
		res.ParentTitle = pMenu.Title
	}

	return
}

// GetAllList 获取所有菜单列表
func (m *Menu) GetAllList(ctx context.Context) (res []*MenuModel, err error) {
	// 获取所有菜单信息
	menus, err := dao.MenuInfo.FindAll(ctx)
	if err != nil {
		return
	}

	// 创建菜单信息映射表
	menuMap := make(map[int64]*entity.MenuInfo)
	for _, menu := range menus {
		menuMap[menu.Id] = menu
	}

	// 转换为MenuModel
	res = make([]*MenuModel, 0, len(menus))
	for _, menu := range menus {
		// 创建菜单模型
		menuModel := &MenuModel{MenuInfo: menu}

		// 获取父菜单名称
		if parentMenu, exists := menuMap[menu.Pid]; exists {
			menuModel.ParentTitle = parentMenu.Title
		}

		// 添加到结果
		res = append(res, menuModel)
	}

	return
}

// Create 创建菜单
func (m *Menu) Create(ctx context.Context, req *MenuCreateModel) (err error) {
	// 创建菜单
	_, err = dao.MenuInfo.Create(ctx, req.MenuInfo)
	return
}

// Update 更新菜单
func (m *Menu) Update(ctx context.Context, req *MenuUpdateModel) (err error) {
	// 更新菜单
	err = dao.MenuInfo.Update(ctx, req.MenuInfo)
	return
}

// Delete 删除菜单
func (m *Menu) Delete(ctx context.Context, id int64) (err error) {
	// 删除菜单
	err = dao.MenuInfo.Delete(ctx, id)
	return
}
