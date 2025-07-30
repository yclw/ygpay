package menu

import (
	"context"

	v2 "yclw/ygpay/api/menu/v2"
	"yclw/ygpay/internal/logic/menu"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV2) Update(ctx context.Context, req *v2.UpdateReq) (res *v2.UpdateRes, err error) {
	err = c.menuService.Update(ctx, c.updateReqToUpdateModel(req))
	if err != nil {
		return
	}
	return
}

func (c *ControllerV2) updateReqToUpdateModel(req *v2.UpdateReq) *menu.MenuUpdateModel {
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
