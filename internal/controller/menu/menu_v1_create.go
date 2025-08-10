package menu

import (
	"context"

	v1 "yclw/ygpay/api/menu/v1"
	"yclw/ygpay/internal/logic/menu"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	_, err = c.menuService.Create(ctx, c.createReqToCreateModel(req))
	return
}

func (c *ControllerV1) createReqToCreateModel(req *v1.CreateReq) *menu.MenuCreateModel {
	return &menu.MenuCreateModel{
		ParentUid: req.ParentUid,
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
	}
}
