package menu

import (
	"context"

	v1 "yclw/ygpay/api/menu/v1"
	"yclw/ygpay/internal/logic/menu"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = c.menuService.Update(ctx, c.updateReqToUpdateModel(req))
	if err != nil {
		return
	}
	return
}

func (c *ControllerV1) updateReqToUpdateModel(req *v1.UpdateReq) *menu.MenuUpdateModel {
	return &menu.MenuUpdateModel{
		MenuInfo: &do.MenuInfo{
			Id:                req.Id,
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
			Id:  req.Id,
			Pid: req.Pid,
		},
	}
}
