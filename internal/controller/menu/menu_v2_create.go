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
			Pid: req.ParentId,
		},
	}
}
