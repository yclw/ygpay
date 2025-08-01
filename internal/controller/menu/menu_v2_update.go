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
			Id:         req.Id,
			Type:       req.Type,
			Name:       req.Name,
			Path:       req.Path,
			Title:      req.Title,
			Icon:       req.Icon,
			Sort:       req.Sort,
			ShowParent: req.ShowParent,
			ShowLink:   req.ShowLink,
			KeepAlive:  req.KeepAlive,
			Redirect:   req.Redirect,
			Component:  req.Component,
			FrameSrc:   req.FrameSrc,
			Url:        req.Url,
			Status:     req.Status,
		},
		MenuTree: &do.MenuTree{
			Id:  req.Id,
			Pid: req.ParentId,
		},
	}
}
