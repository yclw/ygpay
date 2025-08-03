package role

import (
	"context"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/internal/logic/role"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = c.RoleService.Create(ctx, c.createReqToCreateModel(req))
	return
}

func (c *ControllerV1) createReqToCreateModel(req *v1.CreateReq) *role.RoleCreateModel {
	return &role.RoleCreateModel{
		RoleInfo: &do.RoleInfo{
			Pid:    req.ParentId,
			Name:   req.Name,
			Key:    req.Key,
			Remark: req.Remark,
			Sort:   req.Sort,
			Status: req.Status,
		},
	}
}
