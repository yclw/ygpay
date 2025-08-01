package menu

import (
	"context"

	v2 "yclw/ygpay/api/menu/v2"
	"yclw/ygpay/internal/logic/menu"
)

func (c *ControllerV2) GetOne(ctx context.Context, req *v2.GetOneReq) (res *v2.GetOneRes, err error) {
	menu, err := c.menuService.GetOne(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v2.GetOneRes{
		MenuModel: c.menuModelToV2(menu),
	}, nil
}

func (c *ControllerV2) menuModelToV2(menuInfo *menu.MenuModel) *v2.MenuModel {
	return &v2.MenuModel{
		Id:          menuInfo.MenuInfo.Id,
		Type:        menuInfo.MenuInfo.Type,
		Name:        menuInfo.MenuInfo.Name,
		Path:        menuInfo.MenuInfo.Path,
		Title:       menuInfo.MenuInfo.Title,
		Icon:        menuInfo.MenuInfo.Icon,
		Sort:        menuInfo.MenuInfo.Sort,
		ShowParent:  menuInfo.MenuInfo.ShowParent == 1,
		ShowLink:    menuInfo.MenuInfo.ShowLink == 1,
		KeepAlive:   menuInfo.MenuInfo.KeepAlive == 1,
		ParentId:    &menuInfo.TreeNode.Pid,
		ParentTitle: menuInfo.ParentTitle,
		Redirect:    menuInfo.MenuInfo.Redirect,
		Component:   menuInfo.MenuInfo.Component,
		FrameSrc:    menuInfo.MenuInfo.FrameSrc,
		Url:         menuInfo.MenuInfo.Url,
		Status:      menuInfo.MenuInfo.Status,
		CreatedAt:   menuInfo.MenuInfo.CreatedAt,
		UpdatedAt:   menuInfo.MenuInfo.UpdatedAt,
	}
}
