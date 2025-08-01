package menu

import (
	"context"

	v1 "yclw/ygpay/api/menu/v1"
	"yclw/ygpay/internal/logic/menu"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = c.menuService.Create(ctx, c.createReqToCreateModel(req))
	if err != nil {
		return
	}
	return
}

func (c *ControllerV1) createReqToCreateModel(req *v1.CreateReq) *menu.MenuCreateModel {
	return &menu.MenuCreateModel{
		MenuInfo: &do.MenuInfo{
			// Name:              req.Name,
			// Path:              req.Path,
			// Icon:              req.Icon,
			// Title:             req.Title,
			// ShowParent:        req.ShowParent,
			// Component:         req.Component,
			// NoShowingChildren: req.NoShowingChildren,
			// Value:             req.Value,
			// ShowTooltip:       req.ShowTooltip,
			// ParentId:          req.ParentId,
			// Redirect:          req.Redirect,
			// Description:       req.Description,
			// Sort:              req.Sort,
			// Status:            req.Status,
		},
		MenuTree: &do.MenuTree{
			Pid: req.Pid,
		},
	}
}
