package role

import (
	"context"

	v2 "yclw/ygpay/api/role/v2"
	"yclw/ygpay/internal/logic/role"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV2) Update(ctx context.Context, req *v2.UpdateReq) (res *v2.UpdateRes, err error) {
	err = c.RoleService.Update(ctx, c.updateReqToUpdateModel(req))
	if err != nil {
		return
	}
	res = &v2.UpdateRes{}
	return
}

func (c *ControllerV2) updateReqToUpdateModel(req *v2.UpdateReq) *role.RoleUpdateModel {
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
			Pid: req.ParentId,
		},
	}
}
