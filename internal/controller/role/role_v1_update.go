package role

import (
	"context"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/internal/logic/role"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = c.RoleService.Update(ctx, c.updateReqToUpdateModel(req))
	if err != nil {
		return
	}
	return
}

func (c *ControllerV1) updateReqToUpdateModel(req *v1.UpdateReq) *role.RoleUpdateModel {
	return &role.RoleUpdateModel{
		RoleInfo: &do.RoleInfo{
			Id:     req.Id,
			Name:   req.Name,
			Key:    req.Key,
			Remark: req.Remark,
			Sort:   req.Sort,
			Status: req.Status,
		},
		RoleTree: &do.RoleTree{
			Id:  req.Id,
			Pid: req.Pid,
		},
	}
}
