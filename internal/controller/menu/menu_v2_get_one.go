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

func (c *ControllerV2) menuModelToV2(menu *menu.MenuModel) *v2.MenuModel {
	return &v2.MenuModel{
		Id:                menu.MenuInfo.Id,
		Pid:               menu.TreeNode.Pid,
		Name:              menu.MenuInfo.Name,
		Path:              menu.MenuInfo.Path,
		Icon:              menu.MenuInfo.Icon,
		Title:             menu.MenuInfo.Title,
		ShowParent:        menu.MenuInfo.ShowParent == 1,
		Component:         menu.MenuInfo.Component,
		NoShowingChildren: menu.MenuInfo.NoShowingChildren == 1,
		Value:             menu.MenuInfo.Value,
		ShowTooltip:       menu.MenuInfo.ShowTooltip == 1,
		ParentId:          menu.MenuInfo.ParentId,
		Redirect:          menu.MenuInfo.Redirect,
		Description:       menu.MenuInfo.Description,
		Sort:              menu.MenuInfo.Sort,
		Status:            menu.MenuInfo.Status,
		CreatedAt:         menu.MenuInfo.CreatedAt,
		UpdatedAt:         menu.MenuInfo.UpdatedAt,
	}
}
