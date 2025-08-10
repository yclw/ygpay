package menu

import (
	"context"

	v1 "yclw/ygpay/api/menu/v1"
	"yclw/ygpay/internal/logic/menu"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	menu, err := c.menuService.GetOne(ctx, req.MenuUid)
	if err != nil {
		return
	}

	res = &v1.GetOneRes{
		MenuModel: c.menuModelToV1(menu),
	}

	return
}

func (c *ControllerV1) menuModelToV1(menuInfo *menu.MenuModel) *v1.MenuModel {
	res := &v1.MenuModel{
		MenuUid:     menuInfo.MenuInfo.MenuUid,
		Type:        menuInfo.MenuInfo.Type,
		Name:        menuInfo.MenuInfo.Name,
		Path:        menuInfo.MenuInfo.Path,
		Title:       menuInfo.MenuInfo.Title,
		Icon:        menuInfo.MenuInfo.Icon,
		Sort:        menuInfo.MenuInfo.Sort,
		ShowParent:  menuInfo.MenuInfo.ShowParent == 1,
		ShowLink:    menuInfo.MenuInfo.ShowLink == 1,
		KeepAlive:   menuInfo.MenuInfo.KeepAlive == 1,
		ParentId:    menuInfo.MenuInfo.Pid,
		ParentTitle: menuInfo.ParentTitle,
		Redirect:    menuInfo.MenuInfo.Redirect,
		Component:   menuInfo.MenuInfo.Component,
		FrameSrc:    menuInfo.MenuInfo.FrameSrc,
		Url:         menuInfo.MenuInfo.Url,
		Status:      menuInfo.MenuInfo.Status,
		CreatedAt:   menuInfo.MenuInfo.CreatedAt,
		UpdatedAt:   menuInfo.MenuInfo.UpdatedAt,
	}

	return res
}
