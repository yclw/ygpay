package menu

import (
	"context"

	v1 "yclw/ygpay/api/menu/v1"
	"yclw/ygpay/internal/logic/menu"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = c.menuService.Update(ctx, c.updateReqToUpdateModel(req))
	return
}

func (c *ControllerV1) updateReqToUpdateModel(req *v1.UpdateReq) *menu.MenuUpdateModel {
	return &menu.MenuUpdateModel{
		MenuInfo: &do.MenuInfo{
			Id:         req.Id,
			Pid:        req.ParentId,
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
	}
}
