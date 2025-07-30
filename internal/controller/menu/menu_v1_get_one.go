package menu

import (
	"context"

	v1 "yclw/ygpay/api/menu/v1"
	"yclw/ygpay/internal/logic/menu"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	menu, err := c.menuService.GetOne(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetOneRes{
		MenuModel: c.menuModelToV1(menu),
	}, nil
}

func (c *ControllerV1) menuModelToV1(menu *menu.MenuModel) *v1.MenuModel {
	return &v1.MenuModel{
		Id:                menu.MenuInfo.Id,
		Pid:               menu.TreeNode.Pid,
		Name:              menu.MenuInfo.Name,
		Path:              menu.MenuInfo.Path,
		Icon:              menu.MenuInfo.Icon,
		Title:             menu.MenuInfo.Title,
		ShowParent:        menu.MenuInfo.ShowParent,
		Component:         menu.MenuInfo.Component,
		NoShowingChildren: menu.MenuInfo.NoShowingChildren,
		Value:             menu.MenuInfo.Value,
		ShowTooltip:       menu.MenuInfo.ShowTooltip,
		ParentId:          menu.MenuInfo.ParentId,
		Redirect:          menu.MenuInfo.Redirect,
		Description:       menu.MenuInfo.Description,
		Sort:              menu.MenuInfo.Sort,
		Status:            menu.MenuInfo.Status,
		CreatedAt:         menu.MenuInfo.CreatedAt,
		UpdatedAt:         menu.MenuInfo.UpdatedAt,
	}
}
