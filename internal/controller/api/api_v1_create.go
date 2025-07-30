package api

import (
	"context"

	v1 "yclw/ygpay/api/api/v1"
	"yclw/ygpay/internal/logic/api"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = c.ApiService.Create(ctx, c.createReqToCreateModel(req))
	if err != nil {
		return
	}
	return
}

func (c *ControllerV1) createReqToCreateModel(req *v1.CreateReq) *api.ApiCreateModel {
	return &api.ApiCreateModel{
		ApiInfo: &do.ApiInfo{
			Name:        req.Name,
			Path:        req.Path,
			Method:      req.Method,
			GroupName:   req.GroupName,
			Description: req.Description,
			NeedAuth:    req.NeedAuth,
			RateLimit:   req.RateLimit,
			Sort:        req.Sort,
			Status:      req.Status,
		},
	}
}
