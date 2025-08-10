package menu

import (
	"context"
	"yclw/ygpay/internal/consts"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/model/entity"

	"github.com/google/uuid"
)

var MenuService = NewMenu()

type Menu struct {
}

func NewMenu() *Menu {
	return &Menu{}
}

// GetOne 获取单个菜单信息
func (m *Menu) GetOne(ctx context.Context, menuUid string) (res *MenuModel, err error) {
	// 创建菜单模型
	res = &MenuModel{
		MenuInfo: &entity.MenuInfo{},
	}

	// 获取菜单信息
	res.MenuInfo, err = dao.MenuInfo.FindByMenuUid(ctx, menuUid)
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
func (m *Menu) Create(ctx context.Context, req *MenuCreateModel) (id int64, err error) {
	// 生成MenuUid
	req.MenuInfo.MenuUid = uuid.New().String()

	// 获取父菜单ID
	if req.ParentUid == "" {
		req.MenuInfo.Pid = 0
	} else {
		req.MenuInfo.Pid, err = dao.MenuInfo.FindIdByMenuUid(ctx, req.ParentUid)
		if err != nil {
			return
		}
	}

	// 创建菜单
	id, err = dao.MenuInfo.Create(ctx, req.MenuInfo)
	if err != nil {
		return
	}

	// 添加到超级管理员角色
	_, err = dao.RoleMenu.AddRoleMenus(ctx, consts.SuperAdminRoleId, []int64{id})
	return
}

// Update 更新菜单
func (m *Menu) Update(ctx context.Context, req *MenuUpdateModel) (err error) {
	// 获取父菜单ID
	if req.ParentUid == "" {
		req.MenuInfo.Pid = 0
	} else {
		req.MenuInfo.Pid, err = dao.MenuInfo.FindIdByMenuUid(ctx, req.ParentUid)
		if err != nil {
			return
		}
	}

	// 更新菜单
	err = dao.MenuInfo.UpdateByMenuUid(ctx, req.MenuInfo)
	return
}

// Delete 删除菜单
func (m *Menu) Delete(ctx context.Context, menuUid string) (err error) {
	// 删除菜单
	err = dao.MenuInfo.DeleteByMenuUid(ctx, menuUid)
	return
}
