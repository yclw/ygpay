package role

import (
	"context"

	v2 "yclw/ygpay/api/role/v2"
	"yclw/ygpay/internal/logic/role"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV2) Create(ctx context.Context, req *v2.CreateReq) (res *v2.CreateRes, err error) {
	err = c.RoleService.Create(ctx, c.createReqToCreateModel(req))
	if err != nil {
		return
	}
	res = &v2.CreateRes{}
	return
}

func (c *ControllerV2) createReqToCreateModel(req *v2.CreateReq) *role.RoleCreateModel {
	return &role.RoleCreateModel{
		RoleInfo: &do.RoleInfo{
			Name:   req.Name,
			Key:    req.Key,
			Remark: req.Remark,
			Sort:   req.Sort,
			Status: req.Status,
		},
		RoleTree: &do.RoleTree{
			Pid: req.ParentId,
		},
	}
}
