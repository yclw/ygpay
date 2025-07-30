package api

import (
	"context"

	v1 "yclw/ygpay/api/api/v1"
	"yclw/ygpay/internal/logic/api"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = c.ApiService.Update(ctx, c.updateReqToUpdateModel(req))
	if err != nil {
		return
	}
	return
}

func (c *ControllerV1) updateReqToUpdateModel(req *v1.UpdateReq) *api.ApiUpdateModel {
	return &api.ApiUpdateModel{
		Id: req.Id,
		ApiInfo: &do.ApiInfo{
			Id:          req.Id,
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
