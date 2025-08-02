package menu

import (
	"context"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/model/entity"
	"yclw/ygpay/util/tree"
)

var MenuService = NewMenu()

type Menu struct {
}

func NewMenu() *Menu {
	return &Menu{}
}

// GetOne 获取单个菜单信息
func (m *Menu) GetOne(ctx context.Context, id int64) (res *MenuModel, err error) {
	res = &MenuModel{
		TreeNode: &tree.TreeNode{Id: id},
	}

	// 获取菜单信息
	menu, err := dao.MenuInfo.FindByID(ctx, id)
	if err != nil {
		return
	}
	res.MenuInfo = menu

	// 获得pid
	pid, err := dao.MenuTree.FindPidById(ctx, id)
	if err != nil {
		return
	}
	res.Pid = pid

	pMenu, err := dao.MenuInfo.FindByID(ctx, pid)
	if err != nil {
		return
	}
	res.ParentTitle = pMenu.Title

	return
}

// GetAllList 获取所有菜单列表
func (m *Menu) GetAllList(ctx context.Context) (res []*MenuModel, idTree *tree.IdTree, err error) {
	// 获取所有菜单信息
	menus, err := dao.MenuInfo.FindAll(ctx)
	if err != nil {
		return
	}

	// 获取所有菜单id树关系
	Ts, err := dao.MenuTree.FindAll(ctx)
	if err != nil {
		return
	}

	// 构建菜单树
	Ts = append(Ts, tree.T{Id: 0, Pid: -1})
	idTree, err = tree.BuildTree(Ts, 0)
	if err != nil {
		return
	}

	// 创建菜单信息映射表，用于快速查找父菜单标题
	menuMap := make(map[int64]*entity.MenuInfo)
	for _, menu := range menus {
		menuMap[menu.Id] = menu
	}

	// 转换为MenuModel
	res = make([]*MenuModel, 0, len(menus))
	for _, menu := range menus {
		idTree.NodeMap[menu.Id].Data = menu
		treeNode := idTree.NodeMap[menu.Id]

		menuModel := &MenuModel{
			MenuInfo: menu,
			TreeNode: treeNode,
		}
		if menuMap[treeNode.Pid] != nil {
			menuModel.ParentTitle = menuMap[treeNode.Pid].Title
		}
		res = append(res, menuModel)
	}

	return
}

// GetListWithFilter 获取菜单列表
func (m *Menu) GetListWithFilter(ctx context.Context, page, size int, filter *MenuListFilter) (res []*MenuModel, total int, err error) {
	// 构建查询选项
	options := m.buildQueryOptions(filter)

	// 获取筛选后的菜单信息
	menus, total, err := dao.MenuInfo.FindWithPageAndOptions(ctx, page, size, options...)
	if err != nil {
		return
	}

	// 获取所有菜单id树关系
	Ts, err := dao.MenuTree.FindAll(ctx)
	if err != nil {
		return
	}
	pidMap := make(map[int64]int64)
	for _, t := range Ts {
		pidMap[t.Id] = t.Pid
	}

	// 转换为MenuModel
	res = make([]*MenuModel, 0, len(menus))
	for _, menu := range menus {
		pMenu, err := dao.MenuInfo.FindByID(ctx, pidMap[menu.Id])
		if err != nil || pMenu == nil {
			continue
		}

		menuModel := &MenuModel{
			MenuInfo:    menu,
			TreeNode:    &tree.TreeNode{Id: menu.Id, Pid: pidMap[menu.Id]},
			ParentTitle: pMenu.Title,
		}
		res = append(res, menuModel)
	}

	return
}

// buildQueryOptions 构建查询选项
func (a *Menu) buildQueryOptions(filter *MenuListFilter) []dao.QueryOption {
	if filter == nil {
		return nil
	}

	var options []dao.QueryOption
	cols := dao.MenuInfo.Columns()

	// 状态筛选
	if filter.Status != nil {
		options = append(options, dao.Where(cols.Status, *filter.Status))
	}

	// 菜单类型筛选
	if filter.Type != nil {
		options = append(options, dao.Where(cols.Type, *filter.Type))
	}

	// 菜单名称筛选（模糊匹配）
	if filter.Name != "" {
		options = append(options, dao.WhereLike(cols.Name, "%"+filter.Name+"%"))
	}

	// 菜单路径筛选（模糊匹配）
	if filter.Path != "" {
		options = append(options, dao.WhereLike(cols.Path, "%"+filter.Path+"%"))
	}

	// 菜单标题筛选（模糊匹配）
	if filter.Title != "" {
		options = append(options, dao.WhereLike(cols.Title, "%"+filter.Title+"%"))
	}

	// 日期范围筛选
	if filter.StartDate != nil || filter.EndDate != nil {
		options = append(options, dao.WhereBetween(cols.CreatedAt, filter.StartDate, filter.EndDate))
	}

	// 自定义排序
	if filter.SortField != "" {
		if filter.SortDesc {
			options = append(options, dao.OrderDesc(filter.SortField))
		} else {
			options = append(options, dao.OrderAsc(filter.SortField))
		}
	}

	return options
}

// Create 创建菜单
func (m *Menu) Create(ctx context.Context, req *MenuCreateModel) (err error) {
	// 创建菜单
	id, err := dao.MenuInfo.Create(ctx, req.MenuInfo)
	if err != nil {
		return
	}

	req.MenuTree.Id = id

	// 创建菜单树
	_, err = dao.MenuTree.Create(ctx, req.MenuTree)
	if err != nil {
		return
	}

	return
}

// Update 更新菜单
func (m *Menu) Update(ctx context.Context, req *MenuUpdateModel) (err error) {
	// 更新菜单
	err = dao.MenuInfo.Update(ctx, req.MenuInfo)
	if err != nil {
		return
	}

	// 更新菜单树
	err = dao.MenuTree.Update(ctx, req.MenuTree)
	if err != nil {
		return
	}

	return
}

// Delete 删除菜单
func (m *Menu) Delete(ctx context.Context, id int64) (err error) {
	// 删除菜单
	err = dao.MenuInfo.Delete(ctx, id)
	if err != nil {
		return
	}

	// 删除菜单树
	err = dao.MenuTree.Delete(ctx, id)
	if err != nil {
		return
	}

	return
}
