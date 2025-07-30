package menu

import (
	"context"

	v2 "yclw/ygpay/api/menu/v2"
	"yclw/ygpay/internal/logic/menu"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV2) Create(ctx context.Context, req *v2.CreateReq) (res *v2.CreateRes, err error) {
	err = c.menuService.Create(ctx, c.createReqToCreateModel(req))
	if err != nil {
		return
	}
	return
}

func (c *ControllerV2) createReqToCreateModel(req *v2.CreateReq) *menu.MenuCreateModel {
	return &menu.MenuCreateModel{
		MenuInfo: &do.MenuInfo{
			Name:              req.Name,
			Path:              req.Path,
			Icon:              req.Icon,
			Title:             req.Title,
			ShowParent:        req.ShowParent,
			Component:         req.Component,
			NoShowingChildren: req.NoShowingChildren,
			Value:             req.Value,
			ShowTooltip:       req.ShowTooltip,
			ParentId:          req.ParentId,
			Redirect:          req.Redirect,
			Description:       req.Description,
			Sort:              req.Sort,
			Status:            req.Status,
		},
		MenuTree: &do.MenuTree{
			Pid: req.Pid,
		},
	}
}
