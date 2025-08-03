package role

import (
	"context"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/internal/logic/role"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = c.RoleService.Update(ctx, c.updateReqToUpdateModel(req))
	return
}

// updateReqToUpdateModel 将v1.UpdateReq转换为role.RoleUpdateModel
func (c *ControllerV1) updateReqToUpdateModel(req *v1.UpdateReq) *role.RoleUpdateModel {
	return &role.RoleUpdateModel{
		RoleInfo: &do.RoleInfo{
			Id:     req.Id,
			Pid:    req.ParentId,
			Name:   req.Name,
			Key:    req.Key,
			Remark: req.Remark,
			Sort:   req.Sort,
			Status: req.Status,
		},
	}
}
