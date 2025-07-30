package api

import (
	"context"

	v1 "yclw/ygpay/api/api/v1"
	"yclw/ygpay/internal/logic/api"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	model, err := c.ApiService.GetOne(ctx, req.Id)
	if err != nil {
		return
	}
	res = &v1.GetOneRes{
		ApiModel: c.apiModelToV1(model),
	}
	return
}

func (c *ControllerV1) apiModelToV1(model *api.ApiModel) *v1.ApiModel {
	return &v1.ApiModel{
		Id:          model.ApiInfo.Id,
		Name:        model.ApiInfo.Name,
		Path:        model.ApiInfo.Path,
		Method:      model.ApiInfo.Method,
		GroupName:   model.ApiInfo.GroupName,
		Description: model.ApiInfo.Description,
		NeedAuth:    model.ApiInfo.NeedAuth,
		RateLimit:   model.ApiInfo.RateLimit,
		Sort:        model.ApiInfo.Sort,
		Status:      model.ApiInfo.Status,
		CreatedAt:   model.ApiInfo.CreatedAt,
		UpdatedAt:   model.ApiInfo.UpdatedAt,
	}
}
