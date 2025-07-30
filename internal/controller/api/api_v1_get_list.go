package api

import (
	"context"

	v1 "yclw/ygpay/api/api/v1"
	"yclw/ygpay/internal/logic/api"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	// 构建筛选参数
	filter := &api.ApiListFilter{
		Name:      req.Name,
		Path:      req.Path,
		Status:    req.Status,
		GroupName: req.GroupName,
		Method:    req.Method,
		NeedAuth:  req.NeedAuth,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		SortField: req.SortField,
		SortDesc:  req.SortDesc,
	}

	// 调用带筛选的查询方法
	models, total, err := c.ApiService.GetListWithFilter(ctx, req.Page, req.Size, filter)
	if err != nil {
		return
	}

	// 构建响应
	res = &v1.GetListRes{
		Total: total,
	}
	res.List = make([]*v1.ApiModel, 0, len(models))
	for _, model := range models {
		res.List = append(res.List, c.apiModelToV1(model))
	}
	return
}
